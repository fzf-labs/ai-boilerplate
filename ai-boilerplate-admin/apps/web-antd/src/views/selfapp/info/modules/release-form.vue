<script lang="ts" setup>
import type { SelfAppInfo } from '#/api/v1/self-app';
import type {
  CreateSelfAppReleaseReq,
  SelfAppReleaseInfo,
  UpdateSelfAppReleaseReq,
} from '#/api/v1/self-app-release';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createSelfAppRelease,
  getSelfAppReleaseInfo,
  updateSelfAppRelease,
} from '#/api/v1/self-app-release';
import { $t } from '#/locales';

import { useReleaseFormSchema } from './release-data';

interface FormData {
  appInfo: SelfAppInfo;
  releaseInfo?: null | SelfAppReleaseInfo;
}

const emit = defineEmits(['success']);
const formData = ref<FormData>();

const getTitle = computed(() => {
  return formData.value?.releaseInfo?.id
    ? $t('ui.actionTitle.edit', ['版本发布'])
    : $t('ui.actionTitle.create', ['版本发布']);
});

// 处理解析到的安装包信息
const handleParsePackage = async (
  buildNum: number,
  version: string,
  packageSize?: number,
  packageMd5?: string,
  minOsVersion?: string,
) => {
  await formApi.setFieldValue('buildNum', buildNum);
  await formApi.setFieldValue('version', version);
  await formApi.setFieldValue('title', `版本 ${version}`);
  if (packageSize) {
    await formApi.setFieldValue('packageSize', packageSize);
  }
  if (packageMd5) {
    await formApi.setFieldValue('packageMd5', packageMd5);
  }
  if (minOsVersion) {
    await formApi.setFieldValue('minOsVersion', minOsVersion);
  } else {
    console.warn('minOsVersion 为空，未设置');
  }
  message.success('已自动填充版本信息');
};

const [Form, formApi] = useVbenForm({
  layout: 'horizontal',
  schema: useReleaseFormSchema(handleParsePackage),
  showDefaultActions: false,
});

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    try {
      const { valid } = await formApi.validate();
      if (!valid) {
        return;
      }
      modalApi.lock();
      // 提交表单
      const formValues = (await formApi.getValues()) as any;

      // 处理灰度设备列表
      const graySns = formValues.graySns
        ? (formValues.graySns as string)
            .split('\n')
            .map((sn: string) => sn.trim())
            .filter(Boolean)
        : [];

      const submitData = {
        channel: formValues.channel,
        buildNum: formValues.buildNum,
        version: formValues.version,
        updateType: formValues.updateType,
        title: formValues.title,
        changelog: formValues.changelog,
        packageURL: formValues.packageURL,
        packageSize: formValues.packageSize,
        packageMd5: formValues.packageMd5,
        minOsVersion: formValues.minOsVersion,
        grayStrategy: formValues.grayStrategy,
        publishTime: formValues.publishTime,
        status: formValues.status,
        graySns,
        packageName: formData.value?.appInfo.packageName,
      };
      try {
        if (formData.value?.releaseInfo?.id) {
          // 更新
          await updateSelfAppRelease({
            body: {
              id: formData.value.releaseInfo.id,
              ...submitData,
            } as UpdateSelfAppReleaseReq,
          });
          message.success({
            content: $t('ui.actionMessage.updateSuccess'),
            key: 'action_msg',
          });
        } else {
          // 创建
          await createSelfAppRelease({
            body: submitData as CreateSelfAppReleaseReq,
          });
          message.success({
            content: $t('ui.actionMessage.createSuccess'),
            key: 'action_msg',
          });
        }
        emit('success');
        await modalApi.close();
      } finally {
        modalApi.lock(false);
      }
    } catch (error) {
      console.error('Failed to submit form:', error);
      modalApi.lock(false);
      // 表单验证错误会自动显示，这里不需要额外处理
    }
  },
  onOpenChange: async (isOpen: boolean) => {
    if (isOpen) {
      const data = modalApi.getData() as FormData;
      formData.value = data;

      if (data.releaseInfo?.id) {
        // 编辑模式，获取详细信息
        try {
          const res = await getSelfAppReleaseInfo({
            params: { id: data.releaseInfo.id },
          });
          if (res.info) {
            await formApi.setValues({
              ...res.info,
              graySns: res.info.graySns?.join('\n') || '',
              publishTime: res.info.publishTime || '',
            });
          }
        } catch (error) {
          console.error('Failed to get release info:', error);
          message.error('获取版本信息失败');
        }
      } else {
        // 新增模式，设置默认值
        await formApi.setValues({
          packageName: data.appInfo.packageName,
          updateType: 2,
          grayStrategy: 1,
          status: -1,
        });
      }
    }
  },
});

defineExpose({ modalApi });
</script>

<template>
  <Modal
    :title="getTitle"
    class="vben-modal"
    :width="800"
    :destroy-on-close="true"
  >
    <div class="px-6 py-4">
      <Form />
    </div>
  </Modal>
</template>

<style scoped>
:deep(.vben-modal .ant-modal-body) {
  padding: 0;
}
</style>
