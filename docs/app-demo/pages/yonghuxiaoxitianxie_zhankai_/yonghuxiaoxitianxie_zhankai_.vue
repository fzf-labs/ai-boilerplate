<template>
  <view class="flex-col page">
    <view class="flex-col group">
      <image
        class="self-start image"
        src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=967c905996e248889beb79d7e6f3c61e.png"
      />
      <view class="flex-row justify-center items-center self-stretch relative group_2 mt-7">
        <image
          class="image_2 pos"
          src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=06d571df05cc4eb3bbe76d89494b5c8e.png"
        />
        <text class="text">报名消息填写</text>
        <image
          class="image_3 pos_2"
          src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=b78936435301ae54fd85005fc2cba0fc.png"
        />
      </view>
    </view>
    <view class="flex-col mt-27">
      <view class="flex-col self-stretch group_3">
        <view class="flex-row justify-between items-center group_4">
          <view class="flex-col justify-start items-center relative text-wrapper">
            <text class="font text_2">1</text>
          </view>
          <view class="section"></view>
          <view class="flex-col justify-start items-center relative text-wrapper">
            <text class="font text_3">2</text>
          </view>
          <view class="section"></view>
          <view class="flex-col justify-start items-center text-wrapper_2"><text class="font text_4">3</text></view>
        </view>
        <view class="flex-row justify-between mt-7">
          <text class="font_2 text_5">用户报名</text>
          <text class="font_2 text_6">填写信息</text>
          <text class="font_2 text_7">完成报名</text>
        </view>
      </view>
      <view class="flex-col justify-start self-stretch relative group_5">
        <view class="flex-col justify-start items-end relative section_2">
          <view class="section_3"></view>
          <view class="section_4 pos_4"></view>
        </view>
        <image
          class="image_4 pos_3"
          src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=b5dbe8f5b3c587b4076e265b4542d0f8.png"
        />
      </view>
      <view class="flex-col justify-start items-center self-center text-wrapper_3">
        <text class="text_8">提交</text>
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
  .mt-7 {
    margin-top: 14rpx;
  }
  .mt-27 {
    margin-top: 54rpx;
  }
  .page {
    padding-bottom: 52rpx;
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
        margin: 0 26rpx;
        padding: 20rpx 0;
        .image_2 {
          box-shadow: 1rpx 1rpx 4rpx #ffffffb3, -1rpx -1rpx 0rpx #ffffff40 inset, 1rpx 1rpx 0rpx #ffffff inset,
            -4rpx -4rpx 8rpx #ffffff, 4rpx 4rpx 8rpx #3e7ae94d;
          border-radius: 50%;
          width: 76rpx;
          height: 76rpx;
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
    .group_3 {
      padding-left: 60rpx;
      padding-right: 56rpx;
      .group_4 {
        padding: 0 24rpx;
        .text-wrapper {
          padding: 12rpx 0;
          background-color: #3d79e8;
          box-shadow: -0.4rpx -0.4rpx 0rpx #ffffff inset, 0.4rpx 0.4rpx 0rpx #ffffff inset, -2rpx -2rpx 4rpx #ffffff,
            2rpx 2rpx 4rpx #3e7ae933;
          border-radius: 50%;
          width: 46rpx;
          height: 46rpx;
          .text_2 {
            color: #ffffff;
            line-height: 20.52rpx;
          }
          .text_3 {
            color: #f5f5f5;
            line-height: 20.86rpx;
          }
        }
        .section {
          background-color: #3e7ae9;
          width: 184.42rpx;
          height: 1rpx;
        }
        .text-wrapper_2 {
          margin-right: 4rpx;
          padding: 12rpx 0;
          background-color: #f5f5f5;
          box-shadow: -0.4rpx -0.4rpx 0rpx #ffffff inset, 0.4rpx 0.4rpx 0rpx #ffffff inset, -2rpx -2rpx 4rpx #ffffff,
            2rpx 2rpx 4rpx #3e7ae933;
          border-radius: 50%;
          width: 46rpx;
          height: 46rpx;
          .text_4 {
            color: #71828d;
          }
        }
        .font {
          font-size: 28rpx;
          font-family: HarmonyOSSansSC;
          line-height: 21.14rpx;
        }
      }
      .font_2 {
        font-size: 24rpx;
        font-family: HarmonyOSSansSC;
        line-height: 21.14rpx;
        color: #3e7ae9;
      }
      .text_5 {
        line-height: 22.48rpx;
      }
      .text_6 {
        line-height: 22.28rpx;
      }
      .text_7 {
        color: #71828d;
        line-height: 22.6rpx;
      }
    }
    .group_5 {
      margin-top: 44rpx;
      padding-top: 98rpx;
      .section_2 {
        padding: 140rpx 0 620rpx;
        background-image: url('https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=9e3d2e73cc08868097323498b6845195.png');
        background-size: 100% 100%;
        background-repeat: no-repeat;
        .section_3 {
          margin-right: 24rpx;
          background-color: #71828d80;
          width: 1rpx;
          height: 590rpx;
        }
        .section_4 {
          background-color: #70818c80;
          border-radius: 250rpx;
          width: 6rpx;
          height: 620rpx;
        }
        .pos_4 {
          position: absolute;
          right: 22rpx;
          top: 710rpx;
        }
      }
      .image_4 {
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
    .text-wrapper_3 {
      margin-top: 54rpx;
      padding: 36rpx 0;
      background-color: #3d79e8;
      border-radius: 16rpx;
      box-shadow: 1.2rpx 1.2rpx 4rpx #ffffffb3, -1.2rpx -1.2rpx 0rpx #ffffff inset, 1.2rpx 1.2rpx 0rpx #ffffff inset,
        -6rpx -6rpx 10rpx #ffffff, 6rpx 6rpx 10rpx #3e7ae94d;
      width: 636rpx;
      .text_8 {
        color: #f5f5f5;
        font-size: 28rpx;
        font-family: HarmonyOSSansSC;
        line-height: 26.1rpx;
      }
    }
  }
</style>