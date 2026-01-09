<script lang="ts" setup>
import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { FileUpload } from '#/components/upload';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

const emit = defineEmits(['success']);

const fileUrls = ref<string[]>([]);
const fileUploadRef = ref();

const [Form, formApi] = useVbenForm({
  layout: 'horizontal',
  schema: useFormSchema().map((item) => ({ ...item, label: '' })), // 去除label
  showDefaultActions: false,
  commonConfig: {
    hideLabel: true,
  },
});

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }

    // 检查是否有已上传的文件
    if (fileUrls.value.length === 0) {
      message.error('请选择要上传的文件');
      return;
    }

    modalApi.lock();

    try {
      // 文件已经自动上传完成，直接关闭并提示
      await modalApi.close();
      emit('success');
      message.success({
        content: $t('ui.actionMessage.operationSuccess'),
        key: 'action_process_msg',
      });
    } catch (error) {
      console.error('操作失败:', error);
      message.error('操作失败');
    } finally {
      modalApi.lock(false);
    }
  },
});

/** 文件上传变化 */
function handleFileChange(urls: string[]) {
  fileUrls.value = urls;
  formApi.setFieldValue('file', urls);
}
</script>

<template>
  <Modal title="上传文件" ok-text="完成">
    <Form class="mx-4">
      <template #file>
        <div class="w-full">
          <!-- 文件上传组件 -->
          <FileUpload
            ref="fileUploadRef"
            v-model:value="fileUrls"
            :multiple="true"
            :max-number="5"
            :max-size="100"
            :accept="[]"
            :show-description="true"
            @change="handleFileChange"
          />
          <div class="mt-2 text-sm text-gray-500">
            <p>提示：选择文件后将自动开始上传</p>
          </div>
        </div>
      </template>
    </Form>
  </Modal>
</template>
