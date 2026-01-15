<template>
  <view class="flex-col page">
    <view class="flex-col group">
      <image
        class="self-start image"
        src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=967c905996e248889beb79d7e6f3c61e.png"
      />
      <view class="flex-row justify-center items-center self-stretch relative group_2 mt-9">
        <image
          class="image_2 pos"
          src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=3a99794187a0e7a22487e251fc2abd2b.png"
        />
        <text class="text">报名消息填写</text>
        <image
          class="image_3 pos_2"
          src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=b78936435301ae54fd85005fc2cba0fc.png"
        />
      </view>
    </view>
    <view class="mt-28 flex-col">
      <image
        class="self-center image_4"
        src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=d0ce404ac5ac02a23eaf1feeb53f5ee1.png"
      />
      <view class="flex-col self-stretch mt-19">
        <view class="flex-col justify-start self-stretch relative group_3">
          <view class="flex-col items-end section">
            <view class="section_2"></view>
            <view class="section_3"></view>
          </view>
          <image
            class="image_5 pos_3"
            src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=44ce9b556931fbbde2b7be4f9cc298b6.png"
          />
        </view>
        <view class="flex-col justify-start items-center self-center text-wrapper">
          <text class="text_2">提交</text>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
  export default {
    components: {},
    props: {},
    data() {
      return {};
    },

    methods: {},
  };
</script>

<style scoped lang="scss">
  .mt-9 {
    margin-top: 18rpx;
  }
  .mt-19 {
    margin-top: 38rpx;
  }
  .page {
    padding-bottom: 24rpx;
    background-color: #f5f5f5;
    border-radius: 40rpx;
    width: 100%;
    overflow-y: auto;
    overflow-x: hidden;
    height: 100%;
    .group {
      overflow-x: hidden;
      .image {
        margin-left: -32rpx;
        width: 113.0667vw;
        height: 11.7333vw;
      }
      .group_2 {
        margin: 0 28rpx;
        padding: 12rpx 0 20rpx;
        .image_2 {
          box-shadow: 1rpx 1rpx 4rpx #ffffffb3, -1rpx -1rpx 0rpx #ffffff40 inset, 1rpx 1rpx 0rpx #ffffff inset,
            -4rpx -4rpx 8rpx #ffffff, 4rpx 4rpx 8rpx #3e7ae94d;
          border-radius: 50%;
          width: 70rpx;
          height: 70rpx;
        }
        .pos {
          position: absolute;
          left: 0;
          top: 50%;
          transform: translateY(-50%);
        }
        .text {
          color: #3e7ae9;
          font-size: 40rpx;
          font-family: HarmonyOSSansSC;
          line-height: 37.36rpx;
        }
        .image_3 {
          box-shadow: -0.8rpx -0.8rpx 0rpx #ffffff inset, 0.4rpx 0.4rpx 3rpx #ffffff, -6rpx -6rpx 8rpx #ffffff,
            6rpx 6rpx 8rpx #1760e84d;
          border-radius: 284rpx;
          width: 158rpx;
          height: 50rpx;
        }
        .pos_2 {
          position: absolute;
          right: 0;
          top: 50%;
          transform: translateY(-50%);
        }
      }
    }
    .image_4 {
      width: 636rpx;
      height: 88rpx;
    }
    .group_3 {
      padding: 98rpx 0 28rpx;
      .section {
        padding: 140rpx 24rpx 20rpx;
        background-image: url('https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=cc2db9ab13fc67fd566fb869c302ae97.png');
        background-size: 100% 100%;
        background-repeat: no-repeat;
        .section_2 {
          background-color: #70818c80;
          border-radius: 250rpx;
          width: 6rpx;
          height: 620rpx;
        }
        .section_3 {
          background-color: #71828d80;
          width: 1rpx;
          height: 237.9rpx;
        }
      }
      .image_5 {
        overflow: hidden;
        width: 174rpx;
        height: 176rpx;
      }
      .pos_3 {
        position: absolute;
        left: 50%;
        top: 0;
        transform: translateX(-50%);
      }
    }
    .text-wrapper {
      padding: 36rpx 0;
      background-color: #3d79e8;
      border-radius: 16rpx;
      box-shadow: 1.2rpx 1.2rpx 4rpx #ffffffb3, -1.2rpx -1.2rpx 0rpx #ffffff inset, 1.2rpx 1.2rpx 0rpx #ffffff inset,
        -6rpx -6rpx 10rpx #ffffff, 6rpx 6rpx 10rpx #3e7ae94d;
      width: 636rpx;
      .text_2 {
        color: #f5f5f5;
        font-size: 28rpx;
        font-family: HarmonyOSSansSC;
        line-height: 26.1rpx;
      }
    }
  }
</style>