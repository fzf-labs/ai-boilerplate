<script lang="ts" setup>
import type { UploadFile, UploadProps } from 'ant-design-vue';
import type { UploadRequestOption } from 'ant-design-vue/lib/vc-upload/interface';

import type { FileUploadProps } from './typing';

import type { AxiosProgressEvent } from '#/api/infra/file/data';

import { ref, toRefs, watch } from 'vue';

import { IconifyIcon } from '@vben/icons';
import { isFunction, isObject, isString } from '@vben/utils';

import { message, Upload } from 'ant-design-vue';

import { checkFileType } from './helper';
import { UploadResultStatus } from './typing';
import { useUpload, useUploadType } from './use-upload';

defineOptions({ name: 'FileUpload', inheritAttrs: false });

const props = withDefaults(defineProps<FileUploadProps>(), {
  value: () => [],
  disabled: false, // 是否禁用
  helpText: '', // 帮助文本
  maxSize: 2, // MB
  maxNumber: 1, // 最大上传数量
  accept: () => [],
  multiple: false, // 是否多选
  api: undefined, // 上传接口
  resultField: '', // 结果字段
  scene: 'upload', // 场景
  showDescription: false, // 是否显示描述
});
const emit = defineEmits([
  'change',
  'update:value',
  'delete',
  'returnText',
  'fileSelected',
]);
const { accept, helpText, maxNumber, maxSize } = toRefs(props);
const isInnerOperate = ref<boolean>(false);
const { getStringAccept } = useUploadType({
  acceptRef: accept,
  helpTextRef: helpText,
  maxNumberRef: maxNumber,
  maxSizeRef: maxSize,
});

const fileList = ref<UploadProps['fileList']>([]);
const isLtMsg = ref<boolean>(true); // 文件大小错误提示
const isActMsg = ref<boolean>(true); // 文件类型错误提示
const isFirstRender = ref<boolean>(true); // 是否第一次渲染

watch(
  () => props.value,
  (v) => {
    if (isInnerOperate.value) {
      isInnerOperate.value = false;
      return;
    }
    let value: string[] = [];
    if (v) {
      if (Array.isArray(v)) {
        value = v;
      } else {
        value.push(v);
      }
      fileList.value = value.map((item, i) => {
        if (item && isString(item)) {
          return {
            uid: `${-i}`,
            name: item.slice(Math.max(0, item.lastIndexOf('/') + 1)),
            status: UploadResultStatus.DONE,
            url: item,
          };
        } else if (item && isObject(item)) {
          return item;
        }
        return null;
      }) as UploadProps['fileList'];
    }
    if (!isFirstRender.value) {
      emit('change', value);
      isFirstRender.value = false;
    }
  },
  {
    immediate: true,
    deep: true,
  },
);

async function handleRemove(file: UploadFile) {
  if (fileList.value) {
    const index = fileList.value.findIndex((item) => item.uid === file.uid);
    index !== -1 && fileList.value.splice(index, 1);
    const value = getValue();
    isInnerOperate.value = true;
    emit('update:value', value);
    emit('change', value);
    emit('delete', file);
  }
}

async function beforeUpload(file: File) {
  const fileContent = await file.text();
  emit('returnText', fileContent);
  // 发出文件选择事件，传递文件对象
  emit('fileSelected', file);

  const { maxSize, accept } = props;
  const isAct = checkFileType(file, accept);
  if (!isAct) {
    message.error(`请上传${accept.join('/')}格式的文件`);
    isActMsg.value = false;
    // 防止弹出多个错误提示
    setTimeout(() => (isActMsg.value = true), 1000);
  }
  const isLt = file.size / 1024 / 1024 > maxSize;
  if (isLt) {
    message.error(`文件大小不能超过${maxSize}MB`);
    isLtMsg.value = false;
    // 防止弹出多个错误提示
    setTimeout(() => (isLtMsg.value = true), 1000);
  }
  return (isAct && !isLt) || Upload.LIST_IGNORE;
}

async function customRequest(info: UploadRequestOption<any>) {
  let { api, scene } = props;
  if (!api || !isFunction(api)) {
    api = useUpload(scene).httpRequest as any;
  }
  try {
    // 上传文件
    const progressEvent: AxiosProgressEvent = (e: any) => {
      const percent = Math.trunc((e.loaded / e.total!) * 100);
      info.onProgress!({ percent });
    };
    const res = await api?.(info.file as File, progressEvent);
    // 适配OSS上传返回值给ant-design-vue的Upload组件
    const uploadResponse = {
      data: res,
      status: 200,
      statusText: 'OK',
      headers: {},
      config: {},
    };
    info.onSuccess!(uploadResponse);
    // 更新文件
    const value = getValue();
    isInnerOperate.value = true;
    emit('update:value', value);
    emit('change', value);
  } catch (error: any) {
    console.error(error);
    info.onError!(error);
  }
}

function getValue() {
  const list = (fileList.value || [])
    .filter((item) => item?.status === UploadResultStatus.DONE)
    .map((item: any) => {
      if (item?.response && props?.resultField) {
        return item?.response?.data || item?.response;
      }
      return (
        item?.url ||
        item?.response?.data?.url ||
        item?.response?.url ||
        item?.response
      );
    });
  // add by 芋艿：【特殊】单个文件的情况，获取首个元素，保证返回的是 String 类型
  if (props.maxNumber === 1) {
    return list.length > 0 ? list[0] : '';
  }
  return list;
}

function handleChange(info: any) {
  const status = info.file.status;
  if (status === 'done') {
    message.success('上传成功');
  } else if (status === 'error') {
    message.error('上传失败');
  }
}

function handleDrop(_e: DragEvent) {}
</script>

<template>
  <div>
    <Upload.Dragger
      v-bind="$attrs"
      v-model:file-list="fileList"
      :accept="getStringAccept"
      :before-upload="beforeUpload"
      :custom-request="customRequest"
      :disabled="disabled"
      :max-count="maxNumber"
      :multiple="multiple"
      :progress="{ showInfo: true }"
      @change="handleChange"
      @drop="handleDrop"
      @remove="handleRemove"
    >
      <p class="ant-upload-drag-icon flex justify-center">
        <IconifyIcon icon="lucide:inbox" :size="48" />
      </p>
      <p class="ant-upload-text">点击或拖拽文件到此区域上传</p>
      <p class="ant-upload-hint mx-4 mt-2">
        <template v-if="accept.length > 0">
          支持
          <span class="text-primary font-bold">{{ accept.join('/') }}</span>
          格式，
        </template>
        文件大小不超过
        <span class="text-primary font-bold">{{ maxSize }}MB</span>
        <template v-if="maxNumber !== Infinity && maxNumber > 1">
          ，最多可选择
          <span class="text-primary font-bold">{{ maxNumber }}</span> 个文件
        </template>
        <template v-else-if="maxNumber === 1"> ，单个文件上传 </template>
        <template v-if="multiple"> ，支持批量上传 </template>
      </p>
    </Upload.Dragger>
  </div>
</template>
