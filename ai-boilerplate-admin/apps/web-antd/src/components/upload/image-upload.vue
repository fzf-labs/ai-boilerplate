<script lang="ts" setup>
import type { UploadFile, UploadProps } from 'ant-design-vue';
import type { UploadRequestOption } from 'ant-design-vue/lib/vc-upload/interface';

import type { FileUploadProps } from './typing';

import type { AxiosProgressEvent } from '#/api/infra/file/data';

import { ref, toRefs, watch } from 'vue';

import { IconifyIcon } from '@vben/icons';
import { isFunction, isObject, isString } from '@vben/utils';

import { message, Modal, Upload } from 'ant-design-vue';

import { checkImgType, defaultImageAccepts } from './helper';
import { UploadResultStatus } from './typing';
import { useUpload, useUploadType } from './use-upload';

defineOptions({ name: 'ImageUpload', inheritAttrs: false });

const props = withDefaults(defineProps<FileUploadProps>(), {
  value: () => [],
  disabled: false,
  listType: 'picture-card',
  helpText: '',
  maxSize: 2,
  maxNumber: 1,
  accept: () => defaultImageAccepts,
  multiple: false,
  api: undefined,
  resultField: '',
  scene: 'upload',
  showDescription: true,
});
const emit = defineEmits(['change', 'update:value', 'delete']);
const { accept, helpText, maxNumber, maxSize } = toRefs(props);
const isInnerOperate = ref<boolean>(false);
const { getStringAccept } = useUploadType({
  acceptRef: accept,
  helpTextRef: helpText,
  maxNumberRef: maxNumber,
  maxSizeRef: maxSize,
});
const previewOpen = ref<boolean>(false); // 是否展示预览
const previewImage = ref<string>(''); // 预览图片
const previewTitle = ref<string>(''); // 预览标题

const fileList = ref<UploadProps['fileList']>([]);
const isLtMsg = ref<boolean>(true); // 文件大小错误提示
const isActMsg = ref<boolean>(true); // 文件类型错误提示
const isFirstRender = ref<boolean>(true); // 是否第一次渲染

watch(
  () => props.value,
  async (v) => {
    if (isInnerOperate.value) {
      isInnerOperate.value = false;
      return;
    }
    let value: string | string[] = [];
    if (v) {
      value = Array.isArray(v) ? v : [v];
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
    if (isFirstRender.value) {
      isFirstRender.value = false;
    } else {
      emit('change', value);
    }
  },
  {
    immediate: true,
    deep: true,
  },
);

function getBase64<T extends ArrayBuffer | null | string>(file: File) {
  return new Promise<T>((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.addEventListener('load', () => {
      resolve(reader.result as T);
    });
    reader.addEventListener('error', (error) => reject(error));
  });
}

async function handlePreview(file: UploadFile) {
  if (!file.url && !file.preview) {
    file.preview = await getBase64<string>(file.originFileObj!);
  }
  previewImage.value = file.url || file.preview || '';
  previewOpen.value = true;
  previewTitle.value =
    file.name ||
    previewImage.value.slice(
      Math.max(0, previewImage.value.lastIndexOf('/') + 1),
    );
}

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

function handleCancel() {
  previewOpen.value = false;
  previewTitle.value = '';
}

async function beforeUpload(file: File) {
  const { maxSize, accept } = props;
  const isAct = checkImgType(file, accept);
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
    message.success('上传成功');

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
</script>

<template>
  <div>
    <Upload
      v-bind="$attrs"
      v-model:file-list="fileList"
      :accept="getStringAccept"
      :before-upload="beforeUpload"
      :custom-request="customRequest"
      :disabled="disabled"
      :list-type="listType"
      :max-count="maxNumber"
      :multiple="multiple"
      :progress="{ showInfo: true }"
      @preview="handlePreview"
      @remove="handleRemove"
    >
      <div
        v-if="fileList && fileList.length < maxNumber"
        class="flex flex-col items-center justify-center"
      >
        <IconifyIcon icon="lucide:cloud-upload" />
        <div class="mt-2">图片上传</div>
      </div>
    </Upload>
    <div
      v-if="showDescription"
      class="mt-2 flex flex-wrap items-center text-sm"
    >
      请上传不超过
      <div class="text-primary mx-1 font-bold">{{ maxSize }}MB</div>
      的
      <div class="text-primary mx-1 font-bold">{{ accept.join('/') }}</div>
      格式文件
    </div>
    <Modal
      :footer="null"
      :open="previewOpen"
      :title="previewTitle"
      @cancel="handleCancel"
    >
      <img :src="previewImage" alt="" class="w-full" />
    </Modal>
  </div>
</template>

<style>
.ant-upload-select-picture-card {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
