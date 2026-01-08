<!-- 派单管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
    <div class="distribute-page art-full-height">
      <!-- 搜索栏 -->
      <DistributeSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></DistributeSearch>
  
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
        <DistributeView
          v-model="viewDrawerVisible"
          :id="id"
          @submit="refreshData"
        />
        <DistributeSettlementModal
            v-model:visible="settlementModalVisible"
            :id="id"
            @submit="getData"
        />
       
      </ElCard>
    </div>
  </template>
  
<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import { ElTag, ElImage, ElMessageBox } from 'element-plus'
import { useSiteStore } from '@/store/modules/site'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import { DistributeStatus} from '@/enums/statusEnum'
import DistributeSearch from './modules/distribute-search.vue'
import DistributeView from './modules/distribute-view.vue'
import { fetchGetDistributeList, fetchPostDistributeComplete, fetchPostDistributeStart } from '@/api/distribute'
import { DistributeType } from '@/enums/typeEnum'
import { ButtonMoreItem } from '@/components/core/forms/art-button-more/index.vue'
import ArtButtonMore from '@/components/core/forms/art-button-more/index.vue'
import DistributeSettlementModal from './modules/distribute-settlement-modal.vue'


defineOptions({ name: 'Distribute' })
const siteStore = useSiteStore()


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
        apiFn: fetchGetDistributeList,
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
                label: '订单号',
                width: 280,
            },
            {
                prop: 'productInfo',
                label: '商品信息',
                width: 400,
                formatter: (row) => {
                return h('div', { class: 'flex-c' }, [
                    h(ElImage, {
                        class: 'size-12 rounded-md',
                        src: row.product.pic,
                        previewSrcList: [row.product.pic],
                        // 图片预览是否插入至 body 元素上，用于解决表格内部图片预览样式异常
                        previewTeleported: true
                    }),
                    h('div', { class: 'ml-2 flex-1' }, [
                        h('p', { class: 'line-clamp-1' }, row.product.name),
                        h('div', { class: 'flex-1 flex gap-2' }, [
                            h(ElTag, { type:"primary",size:"small" }, () => row.product.category),
                            h(ElTag, { type:"primary",size:"small" }, () => row.product.game)
                        ])
                    ])
                ])
                }
            },
            {
                prop: 'quantity',
                label: '数量',
                formatter: (row) => {
                    return h(ElTag, { type:"primary" }, () => `${row.quantity}${row.unit}`)
                }
            },
            {
                prop: 'type',
                label: '派单类型',
                formatter: (row) => {
                const type = getType(row.type)
                return h(ElTag, { type: type.type }, () => type.text)
                }
            },
            {
                prop: 'totalAmount',
                label: '订单总额',
                formatter: (row) => {
                    return h(ElTag, { type:"primary" }, () => `${row.totalAmount}${siteStore.getInfo.symbol}`)
                }
            },
            {
                prop: 'status',
                label: '服务状态',
                formatter: (row) => {
                const statusConfig = getDistributeStatus(row.status)
                return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
                }
            },
            {
                prop: 'createTime',
                label: '下单时间',
                width: 180,
                sortable: true
            },
            {
                prop: 'operation',
                label: '操作',
                width: 120,
                fixed: 'right', // 固定列
                formatter: (row) =>{
                    const list:ButtonMoreItem[] = []
                    if (row.status == DistributeStatus.Pending) {
                        list.push({
                            key: 'start',
                            label: '开始服务',
                            icon: 'solar:play-circle-bold',
                        })
                    }
                    if (row.status == DistributeStatus.InProgress) {
                        list.push({
                            key: 'completed',
                            label: '完成服务',
                            icon: 'solar:play-circle-bold',
                        })
                    }
                    if (row.status == DistributeStatus.Completed) {
                        list.push({
                            key: 'settlement',
                            label: '报单结算',
                            icon: 'solar:compass-square-bold',
                        })
                    }

                    return h('div', { class: 'distribute flex-c' }, [
                        h(ArtButtonTable, {
                            type: 'view',
                            onClick: () => handleView(row)
                        }),
                        h(ArtButtonMore, {
                            list: list,
                            onClick: (item: ButtonMoreItem) => buttonMoreClick(item, row)
                        })
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

const buttonMoreClick = (item: ButtonMoreItem, row: Distribute.Response.Info) => {
  switch (item.key) {
    case 'start':
    handleStart(row)
    break
    case 'completed':
    handleComplete(row)
    break
    case 'settlement':
    handleSettlement(row)
    break
  }
}


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

const handleStart = (row:Distribute.Response.Info) => {
    ElMessageBox.confirm(`确定要开始服务吗？`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'primary'
    }).then(async() => {
        // TODO: 调用删除接口
        await fetchPostDistributeStart({id:row.id!})
        getData()
    })
    .catch(() => {
        ElMessage.info('已取消')
    })
}

const handleComplete = (row:Distribute.Response.Info) => {
    ElMessageBox.confirm(`确定要服务已完成吗？`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'primary'
    }).then(async() => {
        // TODO: 调用删除接口
        await fetchPostDistributeComplete({id:row.id!})
        getData()
    })
    .catch(() => {
        ElMessage.info('已取消')
    })
}

const handleView = (row:Distribute.Response.Info) => {
    id.value = row.id
    nextTick(() => {
        viewDrawerVisible.value = true
    })
}

const settlementModalVisible = ref(false)
const handleSettlement = (row:Distribute.Response.Info) => {
    id.value = row.id
    nextTick(() => {
        settlementModalVisible.value = true
    })
}

/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Distribute.Response.Info[]): void => {
    selectedRows.value = selection.map((item)=>item.id)
}
</script>
