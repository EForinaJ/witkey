<template>
  <div class="art-card p-5 h-[27.8rem] mb-5 overflow-hidden max-sm:mb-4">
    <div class="art-card-header">
      <div class="title">
        <h4>接单记录</h4>
        <p>近期接单情况</p>
      </div>
      <div @click="goPage" class="flex-cc h-7.5 min-w-17 border border-g-300 rounded-lg text-g-500 c-p">
        <span class="text-xs  mr-1.5">更多</span>
        <ArtSvgIcon icon="solar:multiple-forward-right-bold" class="text-base" />
      </div>
    </div>

    <ElScrollbar style="height: 21.55rem" class="w-full">
      <ArtTable
        :data="tableData!"
        style="margin-top: 0 !important"
        :border="false"
        :stripe="false"
        :header-cell-style="{ background: 'transparent' }"
      >
        <template #default>
          <ElTableColumn label="订单号" prop="code" width="220px"></ElTableColumn>
          <ElTableColumn label="商品信息" prop="product" width="300px">
            <template #default="scope">
              <div class="flex-c">
                <img class="size-12.5 object-cover rounded-md" :src="scope.row.product.pic" />
                <div class="flex flex-col ml-3">
                  <div class="font-medium line-clamp-1">{{ scope.row.product.name }}</div>
                  <div class="flex-1 flex gap-2">
                    <ElTag size="small">{{ scope.row.product.category }}</ElTag>
                    <ElTag size="small">{{ scope.row.product.game }}</ElTag>
                  </div>
                </div>
              </div>
            </template>
          </ElTableColumn>
          <ElTableColumn label="数量" prop="quantity">
            <template #default="scope">
              <ElTag size="small">{{ scope.row.quantity }}{{ scope.row.unit }}</ElTag>
            </template>
          </ElTableColumn>
          <ElTableColumn label="派单类型" prop="type">
            <template #default="scope">
              <ElTag size="small">{{ getType(scope.row.type).text }}</ElTag>
            </template>
          </ElTableColumn>
          <ElTableColumn label="派单状态" prop="type">
            <template #default="scope">
              <ElTag size="small" :type="getDistributeStatus(scope.row.status).type">{{ getDistributeStatus(scope.row.status).text }}</ElTag>
            </template>
          </ElTableColumn>
          <ElTableColumn label="操作" width="60">
            <template #default="scope">
                <div class="flex-c">
                    <ArtButtonTable 
                    @click="handleView(scope.row.id)"
                    type="view"/>
                </div>
            </template>
          </ElTableColumn>
        </template>
      </ArtTable>
    </ElScrollbar>
  </div>
</template>
  
<script setup lang="ts">
import { DistributeStatus } from '@/enums/statusEnum';
import { DistributeType } from '@/enums/typeEnum';

interface Props {
  tableData: Distribute.Response.Info[]
}
const props = withDefaults(defineProps<Props>(), {
})
interface Emits {
  (e: 'view', value: number): void
}
const emit = defineEmits<Emits>()

const router = useRouter()

const TYPE = {
  [DistributeType.Self]: { type: 'primary' as const, text: '个人服务' },
  [DistributeType.Team]: { type: 'primary' as const, text: '自带队伍' },
} as const

const getType = (type: number) => {
  return (
    TYPE[type as keyof typeof TYPE] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}
// 订单状态配置
const DISTRIBUTE_STATUS = {
    [DistributeStatus.Pending]: { type: 'primary' as const, text: '待服务' },
    [DistributeStatus.InProgress]: { type: 'primary' as const, text: '进行中' },
    [DistributeStatus.Completed]: { type: 'success' as const, text: '已完成' },
    [DistributeStatus.Cancel]: { type: 'danger' as const, text: '已取消' },
    [DistributeStatus.Settlementing]: { type: 'warning' as const, text: '结算中' },
    [DistributeStatus.Settlemented]: { type: 'success' as const, text: '已结算' },
} as const

/**
 * 获取订单状态配置
 */
const getDistributeStatus = (status: number) => {
    return (
        DISTRIBUTE_STATUS[status as keyof typeof DISTRIBUTE_STATUS] || {
            type: 'info' as const,
            text: '未知'
        }
    )
}

const handleView = (id:number) => {
  emit('view',id)
}

const goPage = () =>{
  router.push({
    path: '/dashboard/distribute'
  })
}
</script>