<template>
  <div class="art-card p-5 flex-b mb-5 max-sm:mb-4">
    <div>
      <h2 class="text-2xl font-medium">{{ site.title }}</h2>
      <p class="text-g-700 mt-1">各位顶尖打手，欢迎入驻！</p>
      <p class="text-g-700 mt-1">感谢你选择与我们并肩作战。在这里，我们不仅是“带老板”，更是在提供稳定、专业、值得信赖的竞技服务。我们的共同目标，是让每一位客户都获得超预期的体验，从而建立长期信任与口碑。</p>

      <div class="flex flex-wrap gap-3.5 max-w-150 mt-9">
        <div
          class="w-60 flex-cb h-12.5 px-3.5 border border-g-300 c-p rounded-lg text-sm bg-g-100 duration-300 hover:-translate-y-1 max-sm:w-full"
          v-for="link in linkList"
          :key="link.label"
          @click="goPage(link.url)"
        >
          <span class="text-g-700">{{ link.label }}</span>
          <ArtSvgIcon icon="ri:arrow-right-s-line" class="text-lg text-g-600" />
        </div>
      </div>
    </div>
    <img class="w-75 max-md:!hidden" src="@imgs/draw/draw1.png" alt="draw1" />
  </div>
</template>

<script setup lang="ts">

import { useSiteStore } from '@/store/modules/site';
import { WEB_LINKS } from '@/utils/constants'
const {getInfo:site} = useSiteStore()


const linkList = [
  { label: '接单流程', url: WEB_LINKS.DOCS },
  { label: '服务准则', url: WEB_LINKS.GITHUB_HOME },
  { label: '隐私政策', url: WEB_LINKS.INTRODUCE },
  { label: '关于我们', url: WEB_LINKS.BILIBILI }
]

/**
 * 在新标签页中打开指定 URL
 * @param url 要打开的网页地址
 */
const goPage = (url: string): void => {
  window.open(url, '_blank', 'noopener,noreferrer')
}
</script>
