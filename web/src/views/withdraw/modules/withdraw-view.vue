<template>
    <ElDrawer
      v-model="visible"
      title="提现详情"
      size="30%"
      @close="handleClose"
      :close-on-press-escape="false"
    >
    <ElCard v-loading="loading" shadow="never">
        <template #header>
            <div class="flex items-center gap-3">
                <div class="font-bold">
                    提现编号: {{ detail?.code }}
                </div>
                <el-divider direction="vertical" />
                <div class="font-bold">
                    提现时间: {{ detail?.createTime }}
                </div>
            </div>
        </template>
        <ElDescriptions>
            <ElDescriptionsItem :span="3"  label="提现金额">
                <ElTag type="primary">{{ detail?.amount }}{{ site.symbol }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="到账金额">
                <ElTag type="primary">{{ detail?.settledAmount }}{{ site.symbol }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="平台手续费">
                <ElTag type="primary">{{ detail?.serviceFee }}{{ site.symbol }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="提现状态">
                <ElTag :type="getWithdrawStatus(detail?.status!).type">{{ getWithdrawStatus(detail?.status!).text }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="提现类型">
                <ElTag :type="getWithdrawType(detail?.type!).type">{{ getWithdrawType(detail?.type!).text }}</ElTag>
            </ElDescriptionsItem>
        </ElDescriptions>
    </ElCard>

    <ElCard v-if="detail?.status == ApplyStatus.Fail" v-loading="loading" shadow="never" class="mt-[20px]">
        <template #header>
            <div class="font-bold">
                失败原因
            </div>
        </template>
        <div>
            {{ detail?.reason }}
        </div>
    </ElCard>

  </ElDrawer>
</template>
  
<script setup lang="ts">
import { fetchGetWithdrawDetail } from '@/api/withdraw';
import { ApplyStatus } from '@/enums/statusEnum';
import { WithdrawType } from '@/enums/typeEnum';
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
const WITHDRAW_STATUS = {
    [ApplyStatus.Pending]: { type: 'primary' as const, text: '待审核' },
    [ApplyStatus.Success]: { type: 'primary' as const, text: '提现成功' },
    [ApplyStatus.Fail]: { type: 'danger' as const, text: '提现失败' },
} as const

/**
 * 获取订单状态配置
 */
const getWithdrawStatus = (status: number) => {
    return (
        WITHDRAW_STATUS[status as keyof typeof WITHDRAW_STATUS] || {
            type: 'info' as const,
            text: '未知'
        }
    )
}


// 支付方式配置
const WITHDRAW_TYPE = {
  [WithdrawType.AlyPay]: { type: 'primary' as const, text: '支付宝' },
  [WithdrawType.Wechat]: { type: 'success' as const, text: '微信' },
} as const

/**
 * 获取支付方式配置
 */
const getWithdrawType = (mode: number) => {
  return (
    WITHDRAW_TYPE[mode as keyof typeof WITHDRAW_TYPE] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

const {getInfo:site} = useSiteStore()
const detail = ref<Withdraw.Response.Detail>()
const loading = ref<boolean>(false)
const getData = async () =>{
    loading.value = true
    const res = await fetchGetWithdrawDetail({id:props.id!})
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
  