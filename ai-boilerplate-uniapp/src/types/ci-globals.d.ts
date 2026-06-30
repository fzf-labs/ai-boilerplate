import type * as UniApp from '@dcloudio/uni-app'
import type * as Vue from 'vue'
import type useRequest from '../hooks/useRequest'
import type useUpload from '../hooks/useUpload'
import type { useScroll } from '../hooks/useScroll'

declare global {
  const computed: typeof Vue.computed
  const onBeforeMount: typeof Vue.onBeforeMount
  const onBeforeUnmount: typeof Vue.onBeforeUnmount
  const onHide: typeof UniApp.onHide
  const onLaunch: typeof UniApp.onLaunch
  const onLoad: typeof UniApp.onLoad
  const onMounted: typeof Vue.onMounted
  const onReachBottom: typeof UniApp.onReachBottom
  const onReady: typeof UniApp.onReady
  const onShow: typeof UniApp.onShow
  const onUnload: typeof UniApp.onUnload
  const reactive: typeof Vue.reactive
  const ref: typeof Vue.ref
  const useRequest: typeof useRequest
  const useScroll: typeof useScroll
  const useUpload: typeof useUpload
  const watch: typeof Vue.watch
  const watchEffect: typeof Vue.watchEffect

  type Ref<T = any> = Vue.Ref<T>
}

export {}
