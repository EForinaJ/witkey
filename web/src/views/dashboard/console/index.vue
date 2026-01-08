<!-- 工作台页面 -->
<template>
  <div>
    <ElRow :gutter="20">
      <ElCol :sm="24" :md="12" :lg="16">
        <Banner v-loading="loading"
          :compared-to-yesterday="detail?.comparedToYesterday!"
          :today-commission="detail?.TodayCommission!"
        />
        <DistributeList
        @view="handleView"
        
        v-if="detail?.distributeList" v-loading="loading" 
        :table-data="detail?.distributeList!"/>
        <AboutProject />
      </ElCol>
      <ElCol :sm="24" :md="12" :lg="8">
        <WitkeyCard 
          @withdraw="handleWithdraw"
        />
      </ElCol>
    </ElRow>

    <DistributeView 
      v-model="viewDrawerVisible"
      :id="id"
      @submit="getData"
    />
    <WithdrawModal
      v-model:visible="withdrawVisible"
      @submit="getData"
    />
   
  </div>
</template>

<script setup lang="ts">
import { fetchGetDashboardDetail } from '@/api/dashboard';
import AboutProject from './modules/about-project.vue'
import Banner from './modules/banner.vue'
import DistributeList from './modules/distribute-list.vue'
import WitkeyCard from './modules/witkey-card.vue'
import DistributeView from './modules/distribute-view.vue';
import WithdrawModal from './modules/withdraw-modal.vue';
defineOptions({ name: 'Console' })

const viewDrawerVisible = ref(false)
const id = ref<number>(0)

const detail = ref<Dashboard.Response.Detail>()
const loading = ref<boolean>(false)
const getData = async () =>{
    loading.value = true
    const res = await fetchGetDashboardDetail()
    detail.value = res
    loading.value = false
}

const handleView = (e:number) =>{
  id.value = e
  nextTick(() => {
      viewDrawerVisible.value = true
  })
}
const withdrawVisible = ref(false)
const handleWithdraw = () =>{
  nextTick(() => {
    withdrawVisible.value = true
  })
}


onMounted(()=>{
  getData()
})
</script>
