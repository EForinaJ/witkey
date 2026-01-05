<!-- 提现管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
    <div class="withdraw-page art-full-height">
      <!-- 搜索栏 -->
      <WithdrawSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></WithdrawSearch>
  
      <ElCard class="art-table-card" shadow="never">
        <!-- 表格头部 -->
        <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData"></ArtTableHeader>
        <!-- 表格 -->
        <ArtTable
          :loading="loading"
          :data="data"
          :columns="columns"
          :pagination="pagination"
          @selection-change="handleSelectionChange"
          @pagination:size-change="handleSizeChange"
          @pagination:current-change="handleCurrentChange"
        >
        </ArtTable>
  
        <!-- 订单弹窗 -->
        <WithdrawView
          v-model="viewDrawerVisible"
          :id="id"
          @submit="refreshData"
        />
        <!-- <WithdrawSettlementModal
        v-model:visible="settlementModalVisible"
            :id="id"
            @submit="getData"
        /> -->
       
      </ElCard>
    </div>
  </template>
  
<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import { ElTag } from 'element-plus'
import { useSiteStore } from '@/store/modules/site'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import { ApplyStatus} from '@/enums/statusEnum'
import WithdrawSearch from './modules/withdraw-search.vue'
import WithdrawView from './modules/withdraw-view.vue'
import { WithdrawType } from '@/enums/typeEnum'
// import WithdrawSettlementModal from './modules/withdraw-settlement-modal.vue'
import { fetchGetWithdrawList } from '@/api/withdraw'


defineOptions({ name: 'Withdraw' })
const {getInfo:site} = useSiteStore()


// 弹窗相关

const viewDrawerVisible = ref(false)

const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])
// 搜索表单
const searchForm = ref({
    code: undefined,
    status: undefined
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

const {
    columns,
    columnChecks,
    data,
    loading,
    pagination,
    getData,
    searchParams,
    resetSearchParams,
    handleSizeChange,
    handleCurrentChange,
    refreshData
} = useTable({
    // 核心配置
    core: {
        apiFn: fetchGetWithdrawList,
        apiParams:{
            code: "",
            status: 0,
        },
        paginationKey:{
            current: 'page', 
            size: 'limit'
        },
        columnsFactory: () => [
            {
                prop: 'code',
                label: '提现单号',
                width: 280,
            },
            {
                prop: 'amount',
                label: '提现金额',
                formatter: (row) => {
                    return h(ElTag, { type:"primary" }, () => `${row.amount}${site.symbol}`)
                }
            },
            {
                prop: 'settledAmount',
                label: '到账金额',
                formatter: (row) => {
                    return h(ElTag, { type:"primary" }, () => `${row.settledAmount}${site.symbol}`)
                }
            },
            {
                prop: 'serviceFee',
                label: '第三方平台手续费',
                formatter: (row) => {
                    return h(ElTag, { type:"primary" }, () => `${row.serviceFee}${site.symbol}`)
                }
            },
            {
                prop: 'type',
                label: '到账平台',
                formatter: (row) => {
                const type = getWithdrawType(row.type)
                return h(ElTag, { type: type.type }, () => type.text)
                }
            },
            {
                prop: 'status',
                label: '提现状态',
                formatter: (row) => {
                const statusConfig = getWithdrawStatus(row.status)
                return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
                }
            },
            {
                prop: 'createTime',
                label: '申请时间',
                width: 180,
                sortable: true
            },
            {
                prop: 'operation',
                label: '操作',
                width: 80,
                fixed: 'right', // 固定列
                formatter: (row) =>{
                    return h('div', { class: 'withdraw flex-c' }, [
                        h(ArtButtonTable, {
                            type: 'view',
                            onClick: () => handleView(row)
                        }),
                    ])
                }
            }
        ],
    },
    // 数据处理
    transform: {
        responseAdapter: (response) => {
            return {
                records: response.list,
                total: response.total,
            };
        },
    },
})



/**
 * 搜索处理
 * @param params 参数
 */
const handleSearch = (params: Record<string, any>) => {
    console.log(params)
    // 搜索参数赋值
    Object.assign(searchParams, params)
    getData()
}

const handleView = (row:Withdraw.Response.Info) => {
    id.value = row.id
    nextTick(() => {
        viewDrawerVisible.value = true
    })
}

/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Withdraw.Response.Info[]): void => {
    selectedRows.value = selection.map((item)=>item.id)
}
</script>
