<template>
    <div class="art-card p-4 box-border mb-5 max-sm:mb-4">
        <div class="art-card-header mb-2">
           <div class="flex-c gap-2">
                <ArtSvgIcon class="text-2xl text-primary" icon="solar:chat-round-money-bold"/>
                <h2 class="font-bold">佣金</h2>
           </div>
        </div>
        
        <div class="flex-c justify-between">
            <div class="flex-1 flex items-end gap-2 text-primary">
                <span class="font-bold text-4xl">{{ account.witkey?.commission }}</span>
                <span class="text-sm">{{ site.symbol }}</span>
            </div>
            <div class="ml-3">
                <ElButton @click="handleWithdraw" type="primary">提现</ElButton>
            </div>
        </div>
        <div class="mt-3 text-xs text-g-700 w-fit flex-c gap-1 cursor-pointer">
            查看提现记录
            <ArtSvgIcon icon="solar:double-alt-arrow-right-outline"/>
        </div>
    </div>


    <div class="art-card p-4 box-border mb-5 max-sm:mb-4">
        <ElCard shadow="hover">
            <template #header>
                <div class="card-header flex-c justify-between">
                    <h2 class="title font-bold">{{ account.witkey?.name }}</h2>
                    <div>
                        <ElButton size="small" plain>修改昵称</ElButton>
                        <ElButton size="small" plain>修改相册</ElButton>
                    </div>
                </div>
            </template>
            <ElDescriptions border>
                <ElDescriptionsItem :span="3" label="游戏领域">
                    <ElTag type="primary">{{ account.witkey?.game }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3" label="头衔">
                    <ElTag type="primary">{{ account.witkey?.title }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3" label="评分">
                    <ElRate
                        :model-value="account.witkey?.rate!"
                        disabled
                        show-score
                        text-color="#ff9900"
                        score-template="{value} 分"
                    />
                </ElDescriptionsItem>
            </ElDescriptions>
            <template #footer>
                <ElCarousel height="160px" class="rounded">
                    <ElCarouselItem class="rounded" v-for="item in account.witkey?.album" :key="item">
                        <ElImage class="rounded" style="width: 100%; height: 160px" :src="item" />
                    </ElCarouselItem>
                </ElCarousel>
            </template>
        </ElCard>
    </div>
</template>

<script setup lang="ts">
import { useSiteStore } from '@/store/modules/site';
import { useUserStore } from '@/store/modules/user';

const {getInfo:site} = useSiteStore()
const { getUserInfo:account } = useUserStore()

interface Emits {
  (e: 'withdraw'): void
}
const emit = defineEmits<Emits>()
const handleWithdraw = () =>{
    emit("withdraw")
}
</script>