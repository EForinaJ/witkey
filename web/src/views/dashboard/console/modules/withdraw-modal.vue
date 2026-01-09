<template>
    <ElDialog
      title="提现申请"
      width="20%"
      :model-value="visible"
      align-center
      @update:model-value="handleCancel"
      @close="handleClose"
    >
      <ElForm ref="formRef" :model="form" :rules="rules" label-width="auto">
        <ElFormItem prop="amount">
            <ElInputNumber 
            :precision="2"
            style="width: 100%;"
            v-model="form.amount" placeholder="请输入提现金额" controls-position="right"/>
        </ElFormItem>
        <ElFormItem prop="type">
            <ElSelect
                clearable
                style="width: 100%;"
                v-model="form.type"
                placeholder="请选择提现平台"
                >
                <ElOption
                    label="支付宝"
                    :value="WithdrawType.AlyPay"
                />
                <ElOption
                    label="微信"
                    :value="WithdrawType.Wechat"
                />
            </ElSelect>
        </ElFormItem>
        <ElFormItem prop="name">
            <ElInput v-model="form.name" placeholder="请输入提现真实名称" />
        </ElFormItem>
        <ElFormItem prop="number">
            <ElInput v-model="form.number" placeholder="请输入提现账户号码" />
        </ElFormItem>
      </ElForm>
      <template #footer>
        <ElButton @click="handleCancel">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </template>
    </ElDialog>
  </template>
  
<script setup lang="ts">
import { fetchPostWithdrawCreate } from '@/api/withdraw';
import { WithdrawType } from '@/enums/typeEnum';
import { useSiteStore } from '@/store/modules/site';
import type { FormInstance, FormRules } from 'element-plus'


interface Props {
    visible: boolean
}

interface Emits {
    (e: 'update:visible', value: boolean): void
    (e: 'submit'): void
}

const props = withDefaults(defineProps<Props>(), {
    visible: false,
})

const emit = defineEmits<Emits>()
const {getInfo:site} = useSiteStore()
const formRef = ref<FormInstance>()

/**
 * 弹窗显示状态双向绑定
 */
const visible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
})
/**
 * 表单数据
 */
const form = reactive<Withdraw.Params.Modle>({
    amount: 0,
    name: "",
    number: "",
    type: WithdrawType.AlyPay,
})


const rules = reactive<FormRules>({
    amount: [
        { required: true, message: '请输入提现金额', trigger: 'blur' },
    ],
    name: [
        { required: true, message: '请输入提现名称', trigger: 'blur' },
    ],
    number: [
        { required: true, message: '请输入提现账号', trigger: 'blur' },
    ],
    type: [
        { required: true, message: '请选择提现平台', trigger: 'blur' },
    ],
})


/**
 * 监听弹窗打开，初始化表单数据
 */
watch(
    () => props.visible,
    (newVal) => {
    if (newVal) initForm()
    }
)

/**
 * 初始化表单数据
 * 根据弹窗类型填充表单或重置表单
 */
const initForm = async () => {
  
}

/**
 * 关闭弹窗并重置表单
 */
const handleClose = () => {
    formRef.value?.resetFields()
}

/**
 * 取消操作
 */
const handleCancel = (): void => {
    handleClose()
    emit('update:visible', false)
}

/**
 * 提交表单
 * 验证通过后调用接口保存数据
 */
const handleSubmit = async () => {
    if (!formRef.value) return
    try {
        await formRef.value.validate()
        // TODO: 调用新增/编辑接口
       await fetchPostWithdrawCreate(form)

        ElMessage.success('提交成功')
        emit('submit')
        handleClose()
        handleCancel()
    } catch (error) {
        console.log('表单验证失败:', error)
    }
}
</script>