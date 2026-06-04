<script lang="ts" setup>
import type { MailTemplateInfo } from '#/api/v1/mail-template';

import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { sendMailTemplateMsg } from '#/api/v1/mail-template';

import { useSendMailFormSchema } from '../data';

const emit = defineEmits(['success']);

const formData = ref<MailTemplateInfo>();
const parseParams = (params?: string): string[] => {
  if (!params) {
    return [];
  }
  try {
    const parsed = JSON.parse(params);
    if (Array.isArray(parsed)) {
      return parsed.filter((item) => typeof item === 'string');
    }
  } catch {
    // fall through to split
  }
  return params
    .split(',')
    .map((item) => item.trim())
    .filter(Boolean);
};

const [Form, formApi] = useVbenForm({
  layout: 'horizontal',
  showDefaultActions: false,
});

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    modalApi.lock();
    // 构建发送请求
    const values = await formApi.getValues();
    const paramsObj: Record<string, string> = {};
    parseParams(formData.value?.params).forEach((param) => {
      paramsObj[param] = values[`param_${param}`];
    });
    const sendData = {
      id: formData.value?.id || '',
      mail: values.mail,
      params: paramsObj,
    };

    // 提交表单
    try {
      await sendMailTemplateMsg({ body: sendData });
      // 关闭并提示
      await modalApi.close();
      emit('success');
      message.success({
        content: '邮件发送成功',
        key: 'action_process_msg',
      });
    } catch (error) {
      console.error('发送邮件失败', error);
    } finally {
      modalApi.lock(false);
    }
  },
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 获取数据
    const data = modalApi.getData<MailTemplateInfo>();
    if (!data) {
      return;
    }
    formData.value = data;
    // 更新 form schema
    const schema = buildFormSchema();
    formApi.setState({ schema });
    // 设置到 values
    await formApi.setValues({
      content: data.content,
    });
  },
});

/** 动态构建表单 schema */
const buildFormSchema = () => {
  const schema = useSendMailFormSchema();
  parseParams(formData.value?.params).forEach((param) => {
    schema.push({
      fieldName: `param_${param}`,
      label: `参数 ${param}`,
      component: 'Input',
      componentProps: {
        placeholder: `请输入参数 ${param}`,
      },
      rules: 'required',
    });
  });
  return schema;
};
</script>

<template>
  <Modal title="测试发送邮件">
    <Form class="mx-4" />
  </Modal>
</template>
