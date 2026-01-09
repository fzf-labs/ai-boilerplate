<script setup lang="ts" name="account-select">
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

import { useTabs } from '@vben/hooks';

import { Card, Flex, message, Select } from 'ant-design-vue';

import { getWxGzhAccountSelector } from '#/api/v1/wx-gzh-account';

// 类型定义
interface AccountOption {
  label: string;
  value: string;
}

// 组件props
const props = defineProps({
  modelValue: {
    type: String,
    default: undefined,
  },
  placeholder: {
    type: String,
    default: '请选择公众号',
  },
  width: {
    type: String,
    default: '200px',
  },
});

// 组件emits
const emits = defineEmits<{
  change: [value: string | undefined];
  'update:modelValue': [value: string | undefined];
}>();

// 路由和标签页
const { push } = useRouter();
const tabs = useTabs();

// 账号选择相关
const accountOptions = ref<AccountOption[]>([]);

// 获取账号列表
async function getAccountList() {
  try {
    const res = await getWxGzhAccountSelector({});
    if (res.list && res.list.length > 0) {
      accountOptions.value = res.list
        .filter((item) => item.name && item.appId)
        .map((item) => ({
          label: item.name!,
          value: String(item.appId!),
        }));

      // 如果没有选中值且有账号列表，自动选择第一个
      const firstItem = res.list.find((item) => item.name && item.appId);
      if (!props.modelValue && firstItem) {
        const firstAppId = String(firstItem.appId!);
        emits('update:modelValue', firstAppId);
        emits('change', firstAppId);
      }
    } else {
      message.error(
        '未配置公众号，请在【公众号管理 -> 账号管理】菜单，进行配置',
      );
      await push({ name: 'MpAccount' });
      tabs.closeCurrentTab();
    }
  } catch (error) {
    console.error('获取账号列表失败', error);
    message.error('获取账号列表失败');
  }
}

// 账号选择变化处理
const handleAccountChange = (value: any) => {
  if (value !== undefined && value !== null) {
    const stringValue = String(value);
    emits('update:modelValue', stringValue);
    emits('change', stringValue);
  }
};

// 组件挂载时获取账号列表
onMounted(() => {
  getAccountList();
});
</script>

<template>
  <Card size="small" class="account-select-card">
    <Flex justify="flex-start" align="center">
      <div class="gzh-selector-title">
        <label> 公众号</label>
      </div>
      <div class="gzh-selector-select">
        <Select
          :value="modelValue"
          :style="{ width }"
          :placeholder="placeholder"
          @change="handleAccountChange"
        >
          <Select.Option
            v-for="item in accountOptions"
            :key="item.value"
            :value="item.value"
          >
            {{ item.label }}
          </Select.Option>
        </Select>
      </div>
    </Flex>
  </Card>
</template>

<style scoped lang="scss">
.account-select-card {
  margin-bottom: 16px;
}

.gzh-selector-title {
  margin-right: 16px;
  font-size: 16px;
  font-weight: 600;
  line-height: 16px;
  text-align: center;
  letter-spacing: 0.5px;
}

.gzh-selector-select {
  display: flex;
  align-items: center;
  width: 200px;
}
</style>
