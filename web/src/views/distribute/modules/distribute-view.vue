<template>
    <ElDrawer
      v-model="visible"
      title="订单详情"
      size="40%"
      @close="handleClose"
      :close-on-press-escape="false"
    >
    <ElCard v-loading="loading" shadow="never">
        <template #header>
            <div class="flex items-center gap-3">
                <div class="font-bold">
                    订单编号: {{ detail?.code }}
                </div>
                <el-divider direction="vertical" />
                <div class="font-bold">
                    派单时间: {{ detail?.createTime }}
                </div>
            </div>
        </template>
        
        <ElDescriptions>
            <ElDescriptionsItem :span="3"  label="订单总额">
                <ElTag type="primary">{{ detail?.totalAmount }}{{ site.symbol }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="订单数量">
                <ElTag type="primary">{{ detail?.quantity }}{{ detail?.unit }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="服务状态">
                <ElTag :type="getDistributeStatus(detail?.status!).type">{{ getDistributeStatus(detail?.status!).text }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="派单类型">
                <ElTag :type="getType(detail?.type!).type">{{ getType(detail?.type!).text }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="开始时间">
                {{ detail?.startTime == null ? '未开始' : detail?.startTime }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="完成时间">
                {{ detail?.finishTime == null ? '未完成' : detail?.finishTime }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="用户要求">
                {{ detail?.requirements == "" ? '无备注' : detail?.requirements }}
            </ElDescriptionsItem>
        </ElDescriptions>
        <template #footer>
           <div class="flex-c gap-2">
                <ElImage class="size-15 rounded-md"
                    :src="detail?.product.pic"
                    :preview-src-list="[detail?.product.pic!]"
                />
                <div class="flex-1">
                    <h2 class="line-clamp-1">{{ detail?.product.name }}</h2>
                    <div class="flex gap-2 mt-2">
                        <ElTag type="primary">{{ detail?.product.game }}</ElTag>
                        <ElTag type="primary">{{ detail?.product.category }}</ElTag>
                    </div>
                </div>
           </div>
        </template>
    </ElCard>

    <ElCard v-if="detail?.status == DistributeStatus.Cancel" v-loading="loading" shadow="never" class="mt-[20px]">
        <template #header>
            <div class="font-bold">
                取消原因
            </div>
        </template>
        <div>
            {{ detail?.reason }}
        </div>
    </ElCard>
    <ElCard v-if="detail?.status == DistributeStatus.Settlemented || detail?.status == DistributeStatus.Settlementing" v-loading="loading" shadow="never" class="mt-[20px]">
        <template #header>
            <div class="font-bold">
                报单号: {{ detail?.settlement.code }}
            </div>
        </template>
        <ElDescriptions>
            <ElDescriptionsItem :span="3"  label="报单金额">
                {{ detail?.settlement.amount }} {{ site.symbol }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="所得佣金">
                {{ detail?.settlement.commission }}{{ site.symbol }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="平台抽成">
                {{ detail?.settlement.serviceCharge }}{{ site.symbol }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="结算状态">
                <ElTag :type="getSettlementStatus(detail?.settlement.status!).type">{{ getSettlementStatus(detail?.settlement.status!).text }}</ElTag>
            </ElDescriptionsItem>
        </ElDescriptions>
        <ElSpace wrap>
            <div v-for="(item) in detail?.settlement.images" 
                class="w-[100px] h-[100px]">
                <ElImage
                    class="rounded-lg"
                    :src="item"
                    :preview-src-list="[item]"
                    >
                </ElImage>
            </div>
        </ElSpace>
        <template v-if="detail?.settlement.status == ApplyStatus.Fail" #footer>
            <div class="font-bold">
               拒绝原因：{{detail?.settlement.reason}}
            </div>
        </template>
    </ElCard>
  </ElDrawer>
</template>
  
<script setup lang="ts">
import { fetchGetDistributeDetail } from '@/api/distribute';
import { ApplyStatus, DistributeStatus } from '@/enums/statusEnum';
import { DistributeType } from '@/enums/typeEnum';
import { useSiteStore } from '@/store/modules/site';

interface Props {
    modelValue: boolean
    id?: number | null
}
const props = withDefaults(defineProps<Props>(), {
    modelValue: false,
    id: 0
})
interface Emits {
    (e: 'update:modelValue', value: boolean): void
    (e: 'submit'): void
}
const emit = defineEmits<Emits>()
// 对话框显示控制
const visible = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
})

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


const SETTLEMENT_STATUS = {
    [ApplyStatus.Pending]: { type: 'info' as const, text: '待审核' },
    [ApplyStatus.Success]: { type: 'success' as const, text: '已通过' },
    [ApplyStatus.Fail]: { type: 'danger' as const, text: '已拒绝' },
} as const
const getSettlementStatus = (status: number) => {
    return (
        SETTLEMENT_STATUS[status as keyof typeof SETTLEMENT_STATUS] || {
            type: 'info' as const,
            text: '未知'
        }
    )
}

const {getInfo:site} = useSiteStore()
const detail = ref<Distribute.Response.Detail>()
const loading = ref<boolean>(false)
const getData = async () =>{
    loading.value = true
    const res = await fetchGetDistributeDetail({id:props.id!})
    detail.value = res
    loading.value = false
}


/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const init = async () => {
  await getData()
}
/**
   * 监听弹窗打开，初始化表单数据
   */
watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal) init()
  }
)

/**
 * 关闭弹窗并重置表单
 */
const handleClose = () => {
  visible.value = false
  emit('submit')
}
</script>
  