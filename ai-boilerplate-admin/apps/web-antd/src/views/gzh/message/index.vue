<script lang="ts" setup>
import { Page } from '@vben/common-ui';

import { Button } from 'ant-design-vue';
</script>

<template>
  <Page>
    <Button
      type="link"
      target="_blank"
      href="https://github.com/fzf-labs/ai-boilerplate"
    >
      查看仓库文档
    </Button>
    <br />
    <Button
      type="link"
      target="_blank"
      href="https://github.com/fzf-labs/ai-boilerplate/blob/master/docs/technical-decisions.md"
    >
      查看技术选型说明
    </Button>
  </Page>
</template>
