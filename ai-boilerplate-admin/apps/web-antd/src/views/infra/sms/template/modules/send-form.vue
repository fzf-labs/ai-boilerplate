<script lang="ts" setup>
import type {
  SendSmsTemplateMsgReq,
  SmsTemplateInfo,
} from '#/api/v1/sms-template';

import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { sendSmsTemplateMsg } from '#/api/v1/sms-template';

import { useSendSmsFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<SmsTemplateInfo>();

const [Form, formApi] = useVbenForm({
  layout: 'horizontal',
  showDefaultActions: false,
  commonConfig: {
    labelWidth: 120,
  },
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
    if (formData.value?.templateParams) {
      // 转成对象
      const params = JSON.parse(formData.value.templateParams);
      Object.keys(params).forEach((key) => {
        paramsObj[key] = values[`${key}`];
      });
    }
    const sendData: SendSmsTemplateMsgReq = {
      id: formData.value?.id || '',
      phone: values.mobile,
      params: paramsObj,
    };

    // 提交表单
    try {
      await sendSmsTemplateMsg({ body: sendData });
      // 关闭并提示
      await modalApi.close();
      emit('success');
      message.success({
        content: '短信发送成功',
        key: 'action_process_msg',
      });
    } catch (error) {
      console.error('发送短信失败', error);
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
    const data = modalApi.getData<SmsTemplateInfo>();
    if (!data) {
      return;
    }
    formData.value = data;
    // 更新 form schema
    const schema = buildFormSchema();
    formApi.setState({ schema });
    // 设置到 values
    await formApi.setValues({
      content: data.templateContent,
    });
  },
});

/** 动态构建表单 schema */
const buildFormSchema = () => {
  const schema = useSendSmsFormSchema();
  if (formData.value?.templateParams) {
    const params = JSON.parse(formData.value.templateParams);
    Object.keys(params).forEach((param) => {
      schema.push({
        fieldName: `${param}`,
        label: `参数:${param}`,
        component: 'Input',
        componentProps: {
          placeholder: `请输入参数 ${param}`,
        },
        rules: 'required',
      });
    });
  }
  return schema;
};
</script>

<template>
  <Modal title="发送短信">
    <Form class="mx-4" />
  </Modal>
</template>
