<script>
  export default {
    onLaunch: function () {
      console.log('App Launch');
    },
    onShow: function () {
      console.log('App Show');
    },
    onHide: function () {
      console.log('App Hide');
    },
  };
</script>

<style>
    /************************************************************
  ** 请将全局样式拷贝到项目的全局 CSS 文件或者当前页面的顶部 **
  ** 否则页面将无法正常显示                                  **
  ************************************************************/

  html {
    font-size: 16px;
  }

  body {
    margin: 0;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans',
      'Droid Sans', 'Helvetica Neue', 'Microsoft Yahei', sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  view,
  image,
  text {
    box-sizing: border-box;
    flex-shrink: 0;
  }

  #app {
    width: 100vw;
    height: 100vh;
  }

  .flex-row {
    display: flex;
    flex-direction: row;
  }

  .flex-col {
    display: flex;
    flex-direction: column;
  }

  .justify-start {
    justify-content: flex-start;
  }

  .justify-end {
    justify-content: flex-end;
  }

  .justify-center {
    justify-content: center;
  }

  .justify-between {
    justify-content: space-between;
  }

  .justify-around {
    justify-content: space-around;
  }

  .justify-evenly {
    justify-content: space-evenly;
  }

  .items-start {
    align-items: flex-start;
  }

  .items-end {
    align-items: flex-end;
  }

  .items-center {
    align-items: center;
  }

  .items-baseline {
    align-items: baseline;
  }

  .items-stretch {
    align-items: stretch;
  }

  .self-start {
    align-self: flex-start;
  }

  .self-end {
    align-self: flex-end;
  }

  .self-center {
    align-self: center;
  }

  .self-baseline {
    align-self: baseline;
  }

  .self-stretch {
    align-self: stretch;
  }

  .flex-1 {
    flex: 1 1 0%;
  }

  .flex-auto {
    flex: 1 1 auto;
  }

  .grow {
    flex-grow: 1;
  }

  .grow-0 {
    flex-grow: 0;
  }

  .shrink {
    flex-shrink: 1;
  }

  .shrink-0 {
    flex-shrink: 0;
  }

  .relative {
    position: relative;
  }

  .ml-2 {
    margin-left: 0.86rpx;
  }

  .mt-2 {
    margin-top: 0.86rpx;
  }

  .ml-4 {
    margin-left: 1.72rpx;
  }

  .mt-4 {
    margin-top: 1.72rpx;
  }

  .ml-6 {
    margin-left: 2.58rpx;
  }

  .mt-6 {
    margin-top: 2.58rpx;
  }

  .ml-8 {
    margin-left: 3.44rpx;
  }

  .mt-8 {
    margin-top: 3.44rpx;
  }

  .ml-10 {
    margin-left: 4.3rpx;
  }

  .mt-10 {
    margin-top: 4.3rpx;
  }

  .ml-12 {
    margin-left: 5.16rpx;
  }

  .mt-12 {
    margin-top: 5.16rpx;
  }

  .ml-14 {
    margin-left: 6.02rpx;
  }

  .mt-14 {
    margin-top: 6.02rpx;
  }

  .ml-16 {
    margin-left: 6.88rpx;
  }

  .mt-16 {
    margin-top: 6.88rpx;
  }

  .ml-18 {
    margin-left: 7.74rpx;
  }

  .mt-18 {
    margin-top: 7.74rpx;
  }

  .ml-20 {
    margin-left: 8.6rpx;
  }

  .mt-20 {
    margin-top: 8.6rpx;
  }

  .ml-22 {
    margin-left: 9.46rpx;
  }

  .mt-22 {
    margin-top: 9.46rpx;
  }

  .ml-24 {
    margin-left: 10.32rpx;
  }

  .mt-24 {
    margin-top: 10.32rpx;
  }

  .ml-26 {
    margin-left: 11.18rpx;
  }

  .mt-26 {
    margin-top: 11.18rpx;
  }

  .ml-28 {
    margin-left: 12.04rpx;
  }

  .mt-28 {
    margin-top: 12.04rpx;
  }

  .ml-30 {
    margin-left: 12.9rpx;
  }

  .mt-30 {
    margin-top: 12.9rpx;
  }

  .ml-32 {
    margin-left: 13.76rpx;
  }

  .mt-32 {
    margin-top: 13.76rpx;
  }

  .ml-34 {
    margin-left: 14.62rpx;
  }

  .mt-34 {
    margin-top: 14.62rpx;
  }

  .ml-36 {
    margin-left: 15.48rpx;
  }

  .mt-36 {
    margin-top: 15.48rpx;
  }

  .ml-38 {
    margin-left: 16.34rpx;
  }

  .mt-38 {
    margin-top: 16.34rpx;
  }

  .ml-40 {
    margin-left: 17.2rpx;
  }

  .mt-40 {
    margin-top: 17.2rpx;
  }

  .ml-42 {
    margin-left: 18.06rpx;
  }

  .mt-42 {
    margin-top: 18.06rpx;
  }

  .ml-44 {
    margin-left: 18.92rpx;
  }

  .mt-44 {
    margin-top: 18.92rpx;
  }

  .ml-46 {
    margin-left: 19.78rpx;
  }

  .mt-46 {
    margin-top: 19.78rpx;
  }

  .ml-48 {
    margin-left: 20.64rpx;
  }

  .mt-48 {
    margin-top: 20.64rpx;
  }

  .ml-50 {
    margin-left: 21.5rpx;
  }

  .mt-50 {
    margin-top: 21.5rpx;
  }

  .ml-52 {
    margin-left: 22.36rpx;
  }

  .mt-52 {
    margin-top: 22.36rpx;
  }

  .ml-54 {
    margin-left: 23.22rpx;
  }

  .mt-54 {
    margin-top: 23.22rpx;
  }

  .ml-56 {
    margin-left: 24.08rpx;
  }

  .mt-56 {
    margin-top: 24.08rpx;
  }

  .ml-58 {
    margin-left: 24.94rpx;
  }

  .mt-58 {
    margin-top: 24.94rpx;
  }

  .ml-60 {
    margin-left: 25.8rpx;
  }

  .mt-60 {
    margin-top: 25.8rpx;
  }

  .ml-62 {
    margin-left: 26.66rpx;
  }

  .mt-62 {
    margin-top: 26.66rpx;
  }

  .ml-64 {
    margin-left: 27.52rpx;
  }

  .mt-64 {
    margin-top: 27.52rpx;
  }

  .ml-66 {
    margin-left: 28.38rpx;
  }

  .mt-66 {
    margin-top: 28.38rpx;
  }

  .ml-68 {
    margin-left: 29.24rpx;
  }

  .mt-68 {
    margin-top: 29.24rpx;
  }

  .ml-70 {
    margin-left: 30.1rpx;
  }

  .mt-70 {
    margin-top: 30.1rpx;
  }

  .ml-72 {
    margin-left: 30.96rpx;
  }

  .mt-72 {
    margin-top: 30.96rpx;
  }

  .ml-74 {
    margin-left: 31.82rpx;
  }

  .mt-74 {
    margin-top: 31.82rpx;
  }

  .ml-76 {
    margin-left: 32.68rpx;
  }

  .mt-76 {
    margin-top: 32.68rpx;
  }

  .ml-78 {
    margin-left: 33.54rpx;
  }

  .mt-78 {
    margin-top: 33.54rpx;
  }

  .ml-80 {
    margin-left: 34.4rpx;
  }

  .mt-80 {
    margin-top: 34.4rpx;
  }

  .ml-82 {
    margin-left: 35.26rpx;
  }

  .mt-82 {
    margin-top: 35.26rpx;
  }

  .ml-84 {
    margin-left: 36.12rpx;
  }

  .mt-84 {
    margin-top: 36.12rpx;
  }

  .ml-86 {
    margin-left: 36.98rpx;
  }

  .mt-86 {
    margin-top: 36.98rpx;
  }

  .ml-88 {
    margin-left: 37.84rpx;
  }

  .mt-88 {
    margin-top: 37.84rpx;
  }

  .ml-90 {
    margin-left: 38.7rpx;
  }

  .mt-90 {
    margin-top: 38.7rpx;
  }

  .ml-92 {
    margin-left: 39.56rpx;
  }

  .mt-92 {
    margin-top: 39.56rpx;
  }

  .ml-94 {
    margin-left: 40.42rpx;
  }

  .mt-94 {
    margin-top: 40.42rpx;
  }

  .ml-96 {
    margin-left: 41.28rpx;
  }

  .mt-96 {
    margin-top: 41.28rpx;
  }

  .ml-98 {
    margin-left: 42.14rpx;
  }

  .mt-98 {
    margin-top: 42.14rpx;
  }

  .ml-100 {
    margin-left: 43rpx;
  }

  .mt-100 {
    margin-top: 43rpx;
  }
</style>