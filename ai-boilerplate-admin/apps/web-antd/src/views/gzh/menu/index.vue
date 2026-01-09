<script setup lang="ts" name="gzh-menu">
import type { Rule } from 'ant-design-vue/es/form';

import type { PropType } from 'vue';

import { computed, ref, watch } from 'vue';

import { Page } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import {
  Button,
  Col,
  Empty,
  Form,
  FormItem,
  Input,
  message,
  RadioGroup,
  Row,
} from 'ant-design-vue';

import { getWxGzhMenuInfo, storeWxGzhMenu } from '#/api/v1/wx-gzh-menu';

import AccountSelect from '../components/account-select/index.vue';

// 类型定义
interface MenuButton {
  id: number; // 内部使用的ID，用于UI导航
  type: string; // 按钮类型
  name: string; // 按钮名称
  key?: string; // 按钮key
  value?: string; // 按钮值（可选）
  url?: string; // URL（可选）
  appid?: string; // 小程序appid（可选）
  pagepath?: string; // 小程序页面路径（可选）
  news_info?: {
    list: {
      author: string; // 作者
      content_url: string; // 内容URL
      cover_url: string; // 封面URL
      digest: string; // 摘要
      show_cover: number; // 是否显示封面
      source_url: string; // 来源URL
      title: string; // 标题
    }[];
  }; // 新闻信息（可选）
  sub_button: MenuButton[]; // 子按钮数组
}

interface FormData {
  name: string;
  type: string;
  url: string;
  appid: string;
  pagepath: string;
  key?: string;
  value?: string;
}

// 组件props
const props = defineProps({
  isRemote: {
    type: Boolean,
    default: false,
  },
  menuData: {
    type: Object as PropType<{ button: MenuButton[] }>,
    default: () => ({
      button: [] as MenuButton[],
    }),
  },
});
// 提交菜单数据
const emits = defineEmits<{
  submitData: [data: { appId: number; button: MenuButton[] }];
}>();
// 常量定义
const MIN_NAME_LENGTH = 2;
const MAX_NAME_LENGTH = 4;

const select_menu_id = ref(0);
const select_menu_sub_id = ref(0);

// 账号选择相关
const appId = ref<string | undefined>(undefined);

// 菜单类型选项
const menu_type_options = ref([
  { label: '发送消息', value: 'click' },
  { label: '跳转网页', value: 'view' },
  { label: '跳转小程序', value: 'miniprogram' },
]);

// 表单数据
const formData = ref<FormData>({
  name: '',
  type: '',
  url: '',
  appid: '',
  pagepath: '',
});

// 验证规则
const rule = ref<Record<string, Rule[]>>({
  name: [
    { required: true, message: '请输入菜单名称', trigger: 'blur' },
    {
      min: MIN_NAME_LENGTH,
      message: `请输入菜单名称 最低${MIN_NAME_LENGTH}位字符`,
      trigger: 'blur',
    },
    {
      max: MAX_NAME_LENGTH,
      message: `请输入菜单名称 最多${MAX_NAME_LENGTH}位字符`,
      trigger: 'blur',
    },
  ],
  url: [
    { required: true, message: '请输入网页链接', trigger: 'blur' },
    {
      pattern: /^(http|https):\/\//,
      message: '请输入正确的网页链接',
      trigger: 'blur',
    },
  ],
  appid: [
    { required: true, message: '请输入小程序Appid', trigger: 'blur' },
    {
      pattern: /^[a-z0-9]{18}$/i,
      message: '请输入正确的小程序Appid',
      trigger: 'blur',
    },
  ],
  pagepath: [
    { required: true, message: '请输入页面路径', trigger: 'blur' },
    {
      pattern: /^\/pages\/[a-zA-Z0-9]+\/[a-zA-Z0-9]+$/,
      message: '请输入正确的页面路径',
      trigger: 'blur',
    },
  ],
  key: [{ required: true, message: '请输入菜单内容' }],
});

// 菜单数据
const menu_data = ref<{ button: MenuButton[] }>({
  button: [],
});

// 工具函数
const getCurrentMenu = (): MenuButton | undefined => {
  return menu_data.value.button.find(
    (item) => item.id === select_menu_id.value,
  );
};

const getCurrentSubMenu = (): MenuButton | undefined => {
  const menu = getCurrentMenu();
  return menu?.sub_button.find((item) => item.id === select_menu_sub_id.value);
};

const transformBackendToFrontend = (backendMenu: any): MenuButton => {
  const button: MenuButton = {
    id: 0, // 将在调用处设置
    name: backendMenu.name,
    type: backendMenu.type,
    sub_button: [],
  };

  // 设置可选属性
  if (backendMenu.key) button.key = backendMenu.key;
  if (backendMenu.value) button.value = backendMenu.value;
  if (backendMenu.url) button.url = backendMenu.url;
  if (backendMenu.appid) button.appid = backendMenu.appid;
  if (backendMenu.pagepath) button.pagepath = backendMenu.pagepath;
  if (backendMenu.news_info) button.news_info = backendMenu.news_info;

  // 处理子菜单 - API结构是 sub_button.list
  if (backendMenu.sub_button?.list?.length > 0) {
    button.sub_button = backendMenu.sub_button.list.map(
      (subItem: any, subIndex: number) => ({
        ...transformBackendToFrontend(subItem),
        id: subIndex + 1,
      }),
    );
  }

  return button;
};

const transformFrontendToBackend = (frontendMenu: MenuButton) => {
  const menu: any = {
    name: frontendMenu.name,
    type: frontendMenu.type,
  };

  // 设置可选属性
  if (frontendMenu.key) menu.key = frontendMenu.key;
  if (frontendMenu.value) menu.value = frontendMenu.value;
  if (frontendMenu.url) menu.url = frontendMenu.url;
  if (frontendMenu.appid) menu.appid = frontendMenu.appid;
  if (frontendMenu.pagepath) menu.pagepath = frontendMenu.pagepath;
  if (frontendMenu.news_info) menu.news_info = frontendMenu.news_info;

  // 处理子菜单 - 转换为API结构 sub_button.list
  if (frontendMenu.sub_button?.length > 0) {
    menu.sub_button = {
      list: frontendMenu.sub_button.map((element) =>
        transformFrontendToBackend(element),
      ),
    };
  }

  return menu;
};

// 计算属性
const have_sub_menu = computed(() => {
  const menu = getCurrentMenu();
  return menu ? menu.sub_button.length > 0 : false;
});

// 监听props变化
watch(props, (newVal) => {
  if (newVal.menuData) {
    initData();
  }
});

// 重置菜单数据
const resetMenuData = () => {
  menu_data.value = { button: [] };
  select_menu_id.value = 0;
  select_menu_sub_id.value = 0;
  formData.value = {
    name: '',
    type: '',
    url: '',
    appid: '',
    pagepath: '',
  };
};

// 加载菜单数据
const loadMenuData = async (appId: string) => {
  try {
    const res = await getWxGzhMenuInfo({ params: { appId } });
    if (
      res &&
      res.info &&
      res.info.selfmenuInfo &&
      res.info.selfmenuInfo.button
    ) {
      // 将后端菜单数据转换为前端格式
      const buttons = res.info.selfmenuInfo.button.map((item, index) => ({
        ...transformBackendToFrontend(item),
        id: index + 1,
      }));

      menu_data.value.button = buttons;
      // 如果有菜单，选中第一个
      if (menu_data.value.button.length > 0 && menu_data.value.button[0]) {
        select_menu_id.value = menu_data.value.button[0].id;
        await selectMenu(select_menu_id.value);
      }
    }
  } catch (error) {
    console.error('获取菜单数据失败', error);
  }
};

// 监听账号ID变化
watch(appId, async (newVal) => {
  if (newVal) {
    resetMenuData();
    await loadMenuData(newVal);
  }
});

// 重置表单数据
const resetFormData = () => {
  formData.value = {
    name: '',
    type: '',
    url: '',
    appid: '',
    pagepath: '',
    value: '',
  };
};

// 填充表单数据
const fillFormData = (menu: MenuButton) => {
  formData.value.name = menu.name;
  formData.value.type = menu.type;
  formData.value.key = menu.key || '';
  formData.value.value = menu.value || '';
  formData.value.url = menu.url || '';
  formData.value.appid = menu.appid || '';
  formData.value.pagepath = menu.pagepath || '';
};

// 选择菜单
const selectMenu = async (item: number) => {
  if (item === 0) {
    return;
  }

  resetFormData();
  select_menu_id.value = item;
  select_menu_sub_id.value = 0;

  const menu = getCurrentMenu();
  if (!menu) return;

  // 如果没有二级菜单，将form表单的数据设置为一级菜单的数据
  if (menu.sub_button.length === 0) {
    fillFormData(menu);
  } else {
    // 有二级菜单时只设置名称
    formData.value.name = menu.name;
  }
};

// 选择二级菜单
const selectSubMenu = async (item: number) => {
  if (item === 0) {
    return;
  }

  resetFormData();
  select_menu_sub_id.value = item;

  const menu = getCurrentMenu();
  if (!menu) return;

  const sub_menu = menu.sub_button.find((subItem) => subItem.id === item);
  if (!sub_menu) return;

  fillFormData(sub_menu);
};

// 初始化数据
const initData = async () => {
  if (props.isRemote) {
    // 如果是远程数据 将远程数据赋值给menu_data 给每个菜单添加id
    props.menuData.button.forEach((item: any, index: number) => {
      item.id = index + 1;
      if (item.sub_button.length > 0) {
        item.sub_button.forEach((sub_item: any, sub_index: number) => {
          sub_item.id = sub_index + 1;
        });
      }
    });
    menu_data.value.button = props.menuData.button;
  }
  if (menu_data.value.button.length > 0 && menu_data.value.button[0]) {
    select_menu_id.value = menu_data.value.button[0].id;
    await selectMenu(select_menu_id.value);
  }
};

// 更新菜单按钮数据
const updateMenuButton = (menu: MenuButton, formData: FormData) => {
  const updatedButton: MenuButton = {
    id: menu.id,
    name: formData.name,
    type: formData.type,
    sub_button: menu.sub_button,
  };

  // 设置可选字段
  if (formData.key) updatedButton.key = formData.key;
  if (formData.value) updatedButton.value = formData.value;
  if (formData.url) updatedButton.url = formData.url;
  if (formData.appid) updatedButton.appid = formData.appid;
  if (formData.pagepath) updatedButton.pagepath = formData.pagepath;

  return updatedButton;
};

// 更新子菜单按钮数据
const updateSubMenuButton = (
  subMenu: MenuButton,
  formData: FormData,
): MenuButton => {
  const updatedButton: MenuButton = {
    id: subMenu.id,
    name: formData.name,
    type: formData.type,
    sub_button: [],
  };

  // 设置可选字段
  if (formData.key) updatedButton.key = formData.key;
  if (formData.value) updatedButton.value = formData.value;
  if (formData.url) updatedButton.url = formData.url;
  if (formData.appid) updatedButton.appid = formData.appid;
  if (formData.pagepath) updatedButton.pagepath = formData.pagepath;

  return updatedButton;
};

// 保存菜单数据到本地
const saveMenuData = async () => {
  // 如果没有选择菜单，不保存
  if (select_menu_id.value === 0) {
    message.error('请先选择菜单');
    return;
  }

  const menu = getCurrentMenu();
  if (!menu) return;

  // 如果没有选择二级菜单
  if (select_menu_sub_id.value === 0) {
    if (menu.sub_button.length === 0) {
      // 没有二级菜单，更新完整的一级菜单
      const updatedMenu = updateMenuButton(menu, formData.value);
      const menuIndex = menu_data.value.button.findIndex(
        (item) => item.id === select_menu_id.value,
      );
      if (menuIndex !== -1) {
        menu_data.value.button[menuIndex] = updatedMenu;
      }
    } else {
      // 有二级菜单，只更新一级菜单的名称
      menu.name = formData.value.name;
    }
  } else {
    // 更新二级菜单
    const subMenu = getCurrentSubMenu();
    if (!subMenu) return;

    const updatedSubMenu = updateSubMenuButton(subMenu, formData.value);
    const subMenuIndex = menu.sub_button.findIndex(
      (item) => item.id === select_menu_sub_id.value,
    );
    if (subMenuIndex !== -1) {
      menu.sub_button[subMenuIndex] = updatedSubMenu;
    }
  }
};

// 暂存菜单数据到本地
const tempSaveMenuData = async () => {
  // 如果没有选择菜单，不保存
  if (select_menu_id.value === 0) {
    message.error('请先选择菜单');
    return;
  }

  await saveMenuData();
  message.success('暂存成功');
};

// 提交菜单数据到服务器
const submitMenuData = async () => {
  if (menu_data.value.button.length === 0) {
    message.error('请添加菜单');
    return;
  }
  if (!appId.value) {
    message.error('请选择公众号');
    return;
  }

  // 先保存当前编辑的数据到本地
  await saveMenuData();

  // 将前端菜单数据转换为后端格式
  const menus = menu_data.value.button.map((item) =>
    transformFrontendToBackend(item),
  );

  if (!appId.value) {
    message.error('缺少账号信息');
    return;
  }
  try {
    await storeWxGzhMenu({
      body: {
        appId: appId.value,
        selfmenuInfo: {
          button: menus,
        },
      },
    });
    message.success('保存成功');
    emits('submitData', {
      appId: Number(appId.value),
      button: menu_data.value.button,
    });
  } catch (error) {
    console.error('保存菜单失败', error);
    message.error('保存菜单失败');
  }
};

// 添加一级菜单
const addMainMenu = async () => {
  resetFormData();

  if (menu_data.value.button.length >= 3) {
    return;
  }

  const menu_id =
    menu_data.value.button.length === 0
      ? 1
      : Math.max(...menu_data.value.button.map((item) => item.id)) + 1;

  const newMenu: MenuButton = {
    id: menu_id,
    name: `主菜单${menu_id}`,
    type: 'click',
    key: '',
    sub_button: [],
  };

  menu_data.value.button.push(newMenu);
  select_menu_id.value = menu_id;
  await selectMenu(select_menu_id.value);
};

// 添加二级菜单
const addMSubMenu = async () => {
  resetFormData();

  const menu = getCurrentMenu();
  if (!menu || menu.sub_button.length >= 5) {
    return;
  }

  const sub_menu_id =
    menu.sub_button.length === 0
      ? 1
      : Math.max(...menu.sub_button.map((item) => item.id)) + 1;

  const newSubMenu: MenuButton = {
    id: sub_menu_id,
    name: `子菜单${sub_menu_id}`,
    type: 'click',
    key: '',
    sub_button: [],
  };

  menu.sub_button.push(newSubMenu);
  select_menu_sub_id.value = sub_menu_id;
  await selectSubMenu(select_menu_sub_id.value);

  // 清理一级菜单的其他字段，只保留必要字段
  const cleanedMenu: MenuButton = {
    id: menu.id,
    name: menu.name,
    type: 'click', // 有子菜单时类型固定为click
    sub_button: menu.sub_button,
  };

  const menuIndex = menu_data.value.button.findIndex(
    (item) => item.id === menu.id,
  );
  if (menuIndex !== -1) {
    menu_data.value.button[menuIndex] = cleanedMenu;
  }
};

// 往左移动选择菜单
const onLeftMove = async () => {
  const currentIndex = menu_data.value.button.findIndex(
    (item) => item.id === select_menu_id.value,
  );
  if (currentIndex > 0) {
    const previousButton = menu_data.value.button[currentIndex - 1];
    if (previousButton) {
      const previousId = previousButton.id;
      await selectMenu(previousId);
    }
  }
};

// 往右移动选择菜单
const onRightMove = async () => {
  const currentIndex = menu_data.value.button.findIndex(
    (item) => item.id === select_menu_id.value,
  );
  if (currentIndex < menu_data.value.button.length - 1) {
    const nextButton = menu_data.value.button[currentIndex + 1];
    if (nextButton) {
      const nextId = nextButton.id;
      await selectMenu(nextId);
    }
  }
};

// 删除对应id的菜单
const onDelete = async () => {
  const menuIndex = menu_data.value.button.findIndex(
    (item) => item.id === select_menu_id.value,
  );
  if (menuIndex === -1) return;

  menu_data.value.button.splice(menuIndex, 1);

  // 如果还有菜单，选择相邻的菜单
  if (menu_data.value.button.length > 0) {
    let newSelectedIndex = menuIndex;
    // 如果删除的是最后一个，选择前一个
    if (menuIndex >= menu_data.value.button.length) {
      newSelectedIndex = menu_data.value.button.length - 1;
    }

    const selectedButton = menu_data.value.button[newSelectedIndex];
    if (selectedButton) {
      select_menu_id.value = selectedButton.id;
      await selectMenu(select_menu_id.value);
    }
  } else {
    // 没有菜单了，重置选择
    select_menu_id.value = 0;
    select_menu_sub_id.value = 0;
    resetFormData();
  }
};

// 删除对应id的二级菜单
const onDeleteSub = async () => {
  const menu = getCurrentMenu();
  if (!menu) return;

  const subMenuIndex = menu.sub_button.findIndex(
    (item) => item.id === select_menu_sub_id.value,
  );
  if (subMenuIndex === -1) return;

  menu.sub_button.splice(subMenuIndex, 1);

  // 如果删除的是最后一个二级菜单，选择一级菜单
  if (menu.sub_button.length === 0) {
    select_menu_sub_id.value = 0;
    await selectMenu(select_menu_id.value);
  } else {
    // 选择相邻的二级菜单
    let newSelectedIndex = subMenuIndex;
    if (subMenuIndex >= menu.sub_button.length) {
      newSelectedIndex = menu.sub_button.length - 1;
    }

    const selectedSubButton = menu.sub_button[newSelectedIndex];
    if (selectedSubButton) {
      select_menu_sub_id.value = selectedSubButton.id;
      await selectSubMenu(select_menu_sub_id.value);
    }
  }
};

// 账号选择变化处理
const handleAccountChange = (value: any) => {
  if (value !== undefined && value !== null) {
    appId.value = String(value);
    // 可以在这里添加切换账号时的额外逻辑
    // 例如重新加载该账号的菜单数据
  }
};
</script>

<template>
  <Page>
    <!-- 账号选择 -->
    <div class="gzh-account-selector">
      <div class="gzh-account-selector-left">
        <AccountSelect v-model="appId" @change="handleAccountChange" />
      </div>
      <div class="gzh-account-selector-right">
        <Button type="primary" @click="submitMenuData">保存并发布</Button>
      </div>
    </div>
    <!-- 菜单编辑 -->
    <div class="gzh-menu">
      <Row>
        <Col :span="24">
          <div class="container-box">
            <!--      左侧菜单编辑器-->
            <div class="phone-box">
              <div class="menu__preview-hd"></div>
              <div class="menu__preview-bd">
                <div class="menu__preview-bottom">
                  <div class="menu__keyboard">
                    <div class="menu__keyboard-icon"></div>
                  </div>
                  <div class="menu-list">
                    <Button
                      v-if="menu_data.button.length <= 0"
                      type="text"
                      class="menu-add-no-menu"
                      @click="addMainMenu"
                    >
                      <div class="add-no-menu">
                        <div class="menu__add-icon"></div>
                        <span>添加菜单</span>
                      </div>
                    </Button>
                    <div
                      v-for="item in menu_data.button"
                      v-else
                      :key="item.id"
                      :class="`menu-list-item ${select_menu_id === item.id && select_menu_sub_id === 0 ? 'menu-box-shadow' : ''}`"
                    >
                      <div class="menu__preview-line"></div>
                      <div
                        :class="`menu-item ${select_menu_id === item.id && select_menu_sub_id === 0 ? 'menu-item-color' : ''}`"
                        @click="selectMenu(item.id)"
                      >
                        {{ item.name }}
                      </div>
                      <!--                  二级菜单-->
                      <div v-if="select_menu_id === item.id" class="submenu">
                        <!--                    二级-->
                        <div
                          v-for="sub_item in item.sub_button"
                          :key="sub_item.id"
                          :class="`menu-item-sub ${select_menu_sub_id === sub_item.id ? 'menu-box-shadow menu-item-color' : ''} ${select_menu_sub_id !== sub_item.id ? 'menu-item-border-color' : ''}`"
                        >
                          <span
                            style="
                              display: flex;
                              flex: 1;
                              align-items: center;
                              justify-content: center;
                              height: 100%;
                            "
                            @click="selectSubMenu(sub_item.id)"
                          >
                            {{ sub_item.name }}
                          </span>
                          <div class="sub-menu-bar">
                            <div
                              v-if="select_menu_sub_id === sub_item.id"
                              class="sub-menu-bar-box"
                            >
                              <div
                                class="sub-center-bar bar-hover"
                                @click="onDeleteSub"
                              ></div>
                            </div>
                          </div>
                        </div>
                        <!--                    二级增加按钮-->
                        <div
                          v-if="item.sub_button.length < 5"
                          class="add-button-sub"
                        >
                          <Button
                            type="text"
                            title="最多添加5个子菜单"
                            @click="addMSubMenu"
                          >
                            <template #icon>
                              <Plus />
                            </template>
                          </Button>
                        </div>
                        <i
                          :class="`arrow arrow-out ${select_menu_sub_id === 5 && item.sub_button.length === 5 ? 'arrow-out-select' : ''}`"
                        ></i>
                        <i
                          :class="`arrow arrow-in ${select_menu_sub_id === 5 && item.sub_button.length === 5 ? 'arrow-in-select' : ''}`"
                        ></i>
                      </div>
                      <div class="menu-bar">
                        <div
                          v-if="select_menu_id === item.id"
                          :class="`menu-bar-box ${menu_data.button.length > 1 ? 'bar-padding' : ''}`"
                        >
                          <div
                            v-if="
                              menu_data.button[0] &&
                              item.id !== menu_data.button[0].id
                            "
                            class="left-bar bar-hover"
                            @click="onLeftMove"
                          ></div>
                          <div
                            class="center-bar bar-hover"
                            @click="onDelete"
                          ></div>
                          <div
                            v-if="
                              menu_data.button[1] &&
                              item.id !== menu_data.button[1].id &&
                              menu_data.button.length === 2
                            "
                            class="right-bar bar-hover"
                            @click="onRightMove"
                          ></div>
                          <div
                            v-if="
                              menu_data.button.length > 2 &&
                              menu_data.button[2] &&
                              item.id !== menu_data.button[2].id
                            "
                            class="right-bar bar-hover"
                            @click="onRightMove"
                          ></div>
                        </div>
                      </div>
                    </div>
                    <!--                  主菜单添加按钮-->
                    <div
                      v-if="
                        menu_data.button.length < 3 &&
                        menu_data.button.length > 0
                      "
                      class="add-button"
                    >
                      <div class="menu-button-preview-line"></div>
                      <Button type="text" @click="addMainMenu">
                        <template #icon>
                          <Plus />
                        </template>
                      </Button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <!--        右侧编辑器-->
            <div class="right-editor-box">
              <div v-if="menu_data.button.length > 0" class="attribute-box">
                <!--              没有有二级菜单-->
                <div
                  v-if="
                    (have_sub_menu === false && select_menu_id !== 0) ||
                    select_menu_sub_id !== 0
                  "
                  class="custom-no-have-menu-box"
                >
                  <h3 style="padding-bottom: 25px">菜单配置</h3>
                  <Form
                    :model="formData"
                    :rules="rule"
                    :label-col="{ style: { width: '80px' } }"
                    :wrapper-col="{ style: { width: 'calc(100% - 80px)' } }"
                  >
                    <div class="custom-menu-name">
                      <FormItem label="菜单名称" field="name" required>
                        <Input
                          v-model:value="formData.name"
                          placeholder="仅支持中文或数字"
                        />
                        <template #extra>
                          <div>仅支持中英文和数字，字数不超过4个汉字</div>
                        </template>
                      </FormItem>
                    </div>
                    <div class="custom-menu-type">
                      <FormItem label="消息类型" required>
                        <RadioGroup
                          v-model:value="formData.type"
                          :options="menu_type_options"
                          class="ml-4"
                        />
                      </FormItem>
                    </div>
                    <div
                      v-if="formData.type === 'click'"
                      class="custom-menu-content"
                    >
                      <FormItem label="菜单内容" field="key" required>
                        <Input
                          v-model:value="formData.key"
                          placeholder="仅支持中文或数字"
                        />
                        <template #extra>
                          <div>key值为管理平台创建好的功能key</div>
                        </template>
                      </FormItem>
                    </div>
                    <div
                      v-if="formData.type === 'view'"
                      class="custom-menu-content"
                    >
                      <FormItem label="网页链接" field="url" required>
                        <Input
                          v-model:value="formData.url"
                          placeholder="公众号链接"
                        />
                        <template #extra>
                          <div>跳转连接推荐使用安全域名https://</div>
                        </template>
                      </FormItem>
                    </div>
                    <div
                      v-if="formData.type === 'miniprogram'"
                      class="custom-menu-content"
                    >
                      <FormItem label="AppId" field="appid" required>
                        <Input
                          v-model:value="formData.appid"
                          placeholder="小程序AppId"
                        />
                        <template #extra>
                          <div>
                            输入对应的小程序Appid 示例：wxd027d2b162044fd5
                          </div>
                        </template>
                      </FormItem>
                      <FormItem label="页面路径" field="pagepath" required>
                        <Input
                          v-model:value="formData.pagepath"
                          placeholder="页面路径"
                        />
                        <template #extra>
                          <div>
                            输入对应的小程序页面路径 示例：/pages/index/index
                          </div>
                        </template>
                      </FormItem>
                    </div>
                  </Form>
                </div>
                <!--              有二级菜单-->
                <div
                  v-if="
                    have_sub_menu === true &&
                    select_menu_id !== 0 &&
                    select_menu_sub_id === 0
                  "
                  class="custom-have-menu-box"
                >
                  <h3 style="padding-bottom: 25px">菜单配置</h3>
                  <Form :model="formData" :rules="rule">
                    <div class="custom-menu-name">
                      <FormItem label="菜单名称" field="name" required>
                        <Input
                          v-model:value="formData.name"
                          placeholder="仅支持中文或数字"
                        />
                        <template #extra>
                          <div>仅支持中英文和数字，字数不超过4个汉字</div>
                        </template>
                      </FormItem>
                    </div>
                  </Form>
                </div>
              </div>
              <!--            没有菜单-->
              <div v-if="menu_data.button.length === 0" class="attribute-box">
                <Empty
                  description="你未添加自定义菜单，点击左侧添加菜单为公众号创建菜单栏。"
                />
              </div>
              <!--            右侧底部按钮组-->
              <div class="menu__edit-action-bar">
                <div class="menu__edit-action-inner">
                  <div class="menu__edit-button">
                    <Button type="primary" @click="tempSaveMenuData">
                      暂存
                    </Button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </Col>
      </Row>
    </div>
  </Page>
</template>

<style scoped lang="scss">
@use './style.scss';
</style>
