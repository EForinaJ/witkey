<template>
  <ArtBasicBanner
    class="justify-center !h-53 mb-5 max-sm:!pt-8 max-sm:!h-48 max-sm:mb-4"
    :title="`欢迎回来 ${userInfo?.name}`"
    boxStyle="!bg-theme/10"
    titleColor="var(--art-gray-900)"
    :decoration="false"
    :meteorConfig="{
      enabled: true,
      count: 10
    }"
    :buttonConfig="{
      show: false,
      text: ''
    }"
    :imageConfig="{
      src: bannerCover,
      width: '18rem',
      bottom: '-7.5rem'
    }"
    @click="handleBannerClick"
  >
    <div class="flex mt-6">
      <div class="mr-8 pr-8 border-r border-g-400 dark:border-g-300/60">
        <p class="text-3xl">
          <ArtCountTo
            class="number box-title"
            :target="todayCommission!"
            :duration="1500"
            :suffix="site.symbol"
            separator=","
          />
          <ArtSvgIcon icon="ri:arrow-right-up-line" class="text-xl text-success relative -top-2" />
        </p>
        <p class="mt-1 text-sm text-g-700">今日收益额</p>
      </div>
      <div class="mr-8">
        <p class="text-3xl">
          <ArtCountTo class="number box-title" :target="comparedToYesterday!" :duration="1500" suffix="%" />
          <ArtSvgIcon v-if="comparedToYesterday! > 0" icon="ri:arrow-right-up-line" class="text-xl text-success relative -top-2" />
          <ArtSvgIcon v-if="comparedToYesterday! < 0" icon="ri:arrow-right-down-line" class="text-xl text-error relative -top-2" />
        </p>
        <p class="mt-1 text-sm text-g-700">较昨日</p>
      </div>
    </div>
  </ArtBasicBanner>
</template>

<script setup lang="ts">
import bannerCover from '@imgs/login/lf_icon2.webp'
import { useUserStore } from '@/store/modules/user'
import { useSiteStore } from '@/store/modules/site';

interface Props {
  comparedToYesterday: number | null
  todayCommission: number | null
}

const props = withDefaults(defineProps<Props>(), {
  comparedToYesterday: 0,
  todayCommission: 0,
})


const userStore = useUserStore()
const { getInfo:site } = useSiteStore()
/**
 * 获取当前用户信息
 */
const userInfo = computed(() => userStore.getUserInfo)

/**
 * 处理横幅点击事件
 */
const handleBannerClick = (): void => {
  // TODO: 添加横幅点击处理逻辑
}
</script>