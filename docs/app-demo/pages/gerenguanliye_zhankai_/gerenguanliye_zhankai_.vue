<template>
  <view class="flex-col page">
    <view class="flex-col group">
      <view class="flex-col section">
        <image
          class="image"
          src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=b4643f51e91988d2c492b388930c1619.png"
        />
        <view class="flex-row justify-center items-center relative group_2">
          <text class="text_2">我的活动</text>
          <view class="flex-row pos">
            <view class="flex-col justify-start items-center relative group_3">
              <image
                class="image_2"
                src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=34617b619f377eae0103038bfd0971aa.png"
              />
              <view class="flex-col justify-start items-center text-wrapper pos_2"><text class="text">7</text></view>
            </view>
            <image
              class="ml-14 image_2"
              src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=1d135924b49aa4a89e819eca58466375.png"
            />
          </view>
        </view>
      </view>
      <image
        class="shrink-0 image_3"
        src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=ec9d1fd90d3d7ec4512a440eded8f541.png"
      />
      <view class="flex-col justify-start relative group_4">
        <view class="flex-col section_2">
          <view class="flex-row group_5">
            <view class="self-center section_3 view_2"></view>
            <view class="flex-col self-start group_6">
              <view class="flex-col justify-start items-center self-end relative text-wrapper_4 view">
                <text class="font_2">6</text>
              </view>
              <text class="self-center font text_6">待录取</text>
            </view>
            <view class="self-start section_3 view_3"></view>
            <view class="flex-col self-start group_7">
              <view class="flex-col justify-start items-center self-end relative text-wrapper_4">
                <text class="font_2 text_5">7</text>
              </view>
              <text class="self-center font text_7">待评价</text>
            </view>
          </view>
          <view class="section_4"></view>
        </view>
        <view class="flex-col justify-start items-center text-wrapper_2 pos_3">
          <text class="font text_4">待参加</text>
        </view>
        <view class="flex-col justify-start items-center text-wrapper_3 pos_4"><text class="text_3">5</text></view>
      </view>
      <view class="flex-row items-center relative group_8">
        <view class="flex-col flex-1 section_5">
          <view class="flex-col group_9">
            <view class="self-stretch relative section_6"></view>
            <text class="self-center text_8 mt-11">右划查看已录取其他活动</text>
            <view class="self-stretch divider mt-11"></view>
          </view>
          <view class="flex-col group_10">
            <text class="text_9">【早鸟票】她是谁？--艺术作品展览会</text>
            <view class="mt-12 flex-row justify-between items-center group_11">
              <text class="text_10">耿建奕 网易特邀设计师</text>
              <view class="flex-row items-center">
                <view class="shrink-0 section_8"></view>
                <text class="ml-4 font_3 text_11">了解更多</text>
              </view>
            </view>
          </view>
          <view class="flex-col group_12">
            <view class="flex-row">
              <view class="flex-col justify-start items-center relative text-wrapper_5">
                <text class="font_3">查看活动出行需知</text>
              </view>
              <view class="ml-32 flex-col justify-start items-center text-wrapper_5 view_4">
                <text class="font_3 text_12">加入微信活动群</text>
              </view>
            </view>
            <view class="divider_2 mt-23"></view>
            <view class="flex-col group_13 mt-23">
              <view class="flex-col justify-start items-center self-center image-wrapper">
                <image
                  class="image_4"
                  src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=c765751bbec410c70c090ac9b19581f8.png"
                />
              </view>
              <text class="self-center text_13 mt-15">若二维码无法扫描，请联系工作人员</text>
              <view class="flex-col justify-start items-center self-stretch text-wrapper_6 mt-15">
                <text class="font text_14">生成凭证</text>
              </view>
            </view>
          </view>
        </view>
        <view class="shrink-0 section_7"></view>
        <view class="shrink-0 section_9 pos_5"></view>
      </view>
    </view>
    <image
      class="mt-26 image_5"
      src="https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=2aaccaf46fd69436288b8b59e27ae9f8.png"
    />
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
  .mt-11 {
    margin-top: 22rpx;
  }
  .mt-23 {
    margin-top: 46rpx;
  }
  .mt-15 {
    margin-top: 30rpx;
  }
  .page {
    background-color: #f5f5f5;
    border-radius: 40rpx;
    width: 100%;
    overflow-y: auto;
    overflow-x: hidden;
    height: 100%;
    .group {
      height: 1962rpx;
      .section {
        padding-bottom: 204rpx;
        background-image: url('https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=88470bbe7a03e8ea0f2f5cd94206102e.png');
        background-size: 100% 100%;
        .image {
          width: 100vw;
          height: 11.7333vw;
        }
        .group_2 {
          padding: 20rpx 38rpx;
          .text_2 {
            color: #f5f5f5;
            font-size: 40rpx;
            font-family: HarmonyOSSansSC;
            line-height: 37.4rpx;
          }
          .pos {
            position: absolute;
            right: 38rpx;
            top: 50%;
            transform: translateY(-50%);
            .group_3 {
              width: 80rpx;
              .text-wrapper {
                padding: 8rpx 0;
                background-color: #ffc400;
                box-shadow: -0.4rpx -0.4rpx 0rpx #ffefba inset, 0.4rpx 0.4rpx 0rpx #fffdf5 inset, -2rpx -2rpx 4rpx #ffffff,
                  2rpx 2rpx 4rpx #3e7ae933;
                border-radius: 50%;
                width: 28rpx;
                .text {
                  color: #f5f5f5;
                  font-size: 16rpx;
                  font-family: MiSans;
                  line-height: 12.36rpx;
                }
              }
              .pos_2 {
                position: absolute;
                right: 0;
                top: 0;
              }
            }
            .image_2 {
              box-shadow: 0.6rpx 0.6rpx 4rpx #ffffffb3, -0.6rpx -0.6rpx 0rpx #ffffff inset,
                0.6rpx 0.6rpx 0rpx #ffffff inset, -3rpx -3rpx 8rpx #ffffff, 3rpx 3rpx 8rpx #3e7ae94d;
              border-radius: 50%;
              width: 76rpx;
              height: 76rpx;
            }
          }
        }
      }
      .image_3 {
        margin: -182rpx 38rpx 0;
        width: 674rpx;
        height: 476rpx;
      }
      .group_4 {
        margin-top: 20rpx;
        padding-top: 56rpx;
        .section_2 {
          margin: 0 38rpx;
          padding: 0 16rpx 28rpx;
          filter: drop-shadow(1.4rpx 1.4rpx 2rpx #ffffffb3, -6rpx -6rpx 6rpx #ffffff, 6rpx 6rpx 6rpx #3e7ae94d);
          background-image: url('https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=bd77512a649375dcb9f1239548ef4722.png');
          background-size: 100% 100%;
          background-repeat: no-repeat;
          .group_5 {
            padding: 12rpx 20rpx;
            .section_3 {
              opacity: 0.7;
              background-image: repeating-linear-gradient(
                90deg,
                #71828d,
                #71828d 555.58%,
                transparent 555.58%,
                transparent 560.6310000000001%
              );
              background-position: -220rpx 0rpx;
              width: 1rpx;
              height: 57rpx;
            }
            .view_2 {
              margin-left: 176rpx;
            }
            .group_6 {
              width: 251rpx;
              .view {
                margin-right: 68rpx;
              }
              .text_6 {
                line-height: 25.74rpx;
              }
            }
            .view_3 {
              margin-top: 12rpx;
            }
            .group_7 {
              margin-left: 80rpx;
              width: 97rpx;
              .text_7 {
                line-height: 25.9rpx;
              }
            }
            .text-wrapper_4 {
              padding: 8rpx 0;
              background-color: #ffc400;
              box-shadow: -2rpx -2rpx 4rpx #ffffff, 2rpx 2rpx 4rpx #3e7ae933;
              border-radius: 50%;
              width: 24rpx;
              .font_2 {
                font-size: 16rpx;
                font-family: HarmonyOSSansSC;
                line-height: 11.9rpx;
                color: #ffffff;
              }
              .text_5 {
                line-height: 11.72rpx;
              }
            }
          }
          .section_4 {
            background-color: #f5f5f5;
            border-radius: 524rpx;
            box-shadow: 0rpx 1rpx 0rpx #8f8f8f4d inset, 0rpx -1rpx 0rpx #ffffff inset, 0rpx -4rpx 8rpx #ffffff inset,
              0rpx 4rpx 8rpx #8f8f8f4d inset, -4rpx -4rpx 8rpx #ffffff, 4rpx 4rpx 8rpx #3e7ae94d;
            height: 24rpx;
          }
        }
        .text-wrapper_2 {
          padding: 44rpx 0;
          background-color: #f5f5f5;
          box-shadow: -1rpx -1rpx 0rpx #ffffff inset, 1rpx 1rpx 0rpx #ffffff inset, 3rpx 3rpx 6rpx #ffffff,
            4rpx 4rpx 8rpx #3e7ae966;
          border-radius: 50%;
          width: 114rpx;
          .text_4 {
            color: #3e7ae9;
          }
        }
        .pos_3 {
          position: absolute;
          left: 80rpx;
          top: 0;
        }
        .text-wrapper_3 {
          padding: 8rpx 0;
          background-color: #ffc400;
          box-shadow: -2rpx -2rpx 4rpx #ffffff, 2rpx 2rpx 4rpx #3e7ae933;
          border-radius: 50%;
          width: 32rpx;
          .text_3 {
            color: #ffffff;
            font-size: 20rpx;
            font-family: HarmonyOSSansSC;
            line-height: 14.9rpx;
          }
        }
        .pos_4 {
          position: absolute;
          left: 164rpx;
          top: 8rpx;
        }
      }
      .group_8 {
        margin-top: -38rpx;
        padding-left: 68rpx;
        padding-right: 40rpx;
        .section_5 {
          padding: 0 24rpx 52rpx 30rpx;
          filter: drop-shadow(0rpx -1rpx 1rpx #303030b3, -10rpx -10rpx 9rpx #ffffff, 10rpx 10rpx 9rpx #3e7ae94d);
          background-image: url('https://ide.code.fun/api/image?token=6968f4d34683ce0011d8c657&name=fee591e65793b7d3a0e11203f877d3d3.png');
          background-size: 100% 100%;
          background-repeat: no-repeat;
          height: 1114rpx;
          .group_9 {
            padding-top: 40rpx;
            .section_6 {
              margin-right: 8rpx;
              background-color: #f5f5f5;
              border-radius: 4rpx;
              box-shadow: 0rpx -4rpx 8rpx #ffffff inset, 0rpx 4rpx 8rpx #3e7ae91a inset, 0rpx 4rpx 8rpx #3e7ae933;
              height: 24rpx;
            }
            .text_8 {
              color: #3e7ae9;
              font-size: 20rpx;
              font-family: HarmonyOSSansSC;
              line-height: 18.5rpx;
            }
            .divider {
              margin-left: 16rpx;
              margin-right: 24rpx;
              opacity: 0.9;
              background-image: repeating-linear-gradient(
                90deg,
                #71828d,
                #71828d 0.77%,
                transparent 0.77%,
                transparent 1.5390000000000001%
              );
              background-position: -2rpx 0rpx;
              height: 1.4rpx;
            }
          }
          .group_10 {
            margin-top: 48rpx;
            padding-left: 16rpx;
            .text_9 {
              color: #1e2e39;
              font-size: 32rpx;
              font-family: HarmonyOSSansSC;
              line-height: 30.28rpx;
            }
            .group_11 {
              margin-right: 24rpx;
              .text_10 {
                color: #71828d;
                font-size: 24rpx;
                font-family: HarmonyOSSansSC;
                line-height: 28rpx;
                width: 72rpx;
              }
              .section_8 {
                background-color: #3e7ae9;
                width: 81rpx;
                height: 1rpx;
              }
              .text_11 {
                color: #3e7ae9;
                line-height: 18.68rpx;
              }
            }
          }
          .group_12 {
            margin-top: 232rpx;
            padding: 0 14rpx;
            .text-wrapper_5 {
              flex: 1 1 230rpx;
              padding: 20rpx 0;
              background-color: #f5f5f5;
              border-radius: 8rpx;
              box-shadow: 1rpx 1rpx 4rpx #ffffffb3, -1rpx -1rpx 0rpx #ffffff inset, 1rpx 1rpx 0rpx #ffffff inset,
                -4rpx -4rpx 10rpx #ffffff, 4rpx 4rpx 10rpx #3e7ae94d;
              height: 56rpx;
              .text_12 {
                line-height: 18.6rpx;
              }
            }
            .view_4 {
              margin-right: 8rpx;
            }
            .divider_2 {
              margin: 0 24rpx;
              opacity: 0.9;
              background-image: repeating-linear-gradient(
                90deg,
                #71828d,
                #71828d 0.82%,
                transparent 0.82%,
                transparent 1.638%
              );
              background-position: -2rpx 0rpx;
              height: 3.4rpx;
            }
            .group_13 {
              padding: 0 78rpx;
              .image-wrapper {
                padding: 48rpx 0;
                background-color: #f5f5f5;
                border-radius: 12rpx;
                box-shadow: 1.2rpx 1.2rpx 4rpx #ffffffb3, -1.2rpx -1.2rpx 0rpx #ffffff inset,
                  1.2rpx 1.2rpx 0rpx #ffffff inset, -6rpx -6rpx 14rpx #ffffff, 6rpx 6rpx 14rpx #3e7ae94d;
                width: 296rpx;
                .image_4 {
                  width: 196rpx;
                  height: 196rpx;
                }
              }
              .text_13 {
                color: #a6a6a6;
                font-size: 16rpx;
                font-family: HarmonyOSSansSC;
                font-weight: 300;
                line-height: 15.74rpx;
              }
              .text-wrapper_6 {
                padding: 16rpx 0;
                background-color: #f5f5f5;
                border-radius: 8rpx;
                box-shadow: 1.2rpx 1.2rpx 4rpx #ffffffb3, -1.2rpx -1.2rpx 0rpx #ffffff, 1.2rpx 1.2rpx 0rpx #ffffff inset,
                  -4rpx -4rpx 10rpx #ffffff, 4rpx 4rpx 10rpx #3e7ae94d;
                .text_14 {
                  line-height: 26.26rpx;
                }
              }
            }
          }
          .font_3 {
            font-size: 20rpx;
            font-family: HarmonyOSSansSC;
            line-height: 18.5rpx;
            font-weight: 300;
            color: #3d5464;
          }
        }
        .section_7 {
          margin-left: 24rpx;
          background-color: #71828d80;
          width: 3rpx;
          height: 966rpx;
        }
        .section_9 {
          background-color: #70818c80;
          border-radius: 156rpx;
          width: 8rpx;
          height: 100rpx;
        }
        .pos_5 {
          position: absolute;
          right: 38rpx;
          top: 358rpx;
        }
      }
      .font {
        font-size: 28rpx;
        font-family: HarmonyOSSansSC;
        line-height: 26.58rpx;
        color: #3d5464;
      }
    }
    .image_5 {
      width: 100vw;
      height: 25.3333vw;
    }
  }
</style>