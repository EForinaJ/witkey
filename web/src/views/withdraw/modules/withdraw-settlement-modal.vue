<template>
    <ElDialog
      title="报单结算"
      width="25%"
      :model-value="visible"
      align-center
      @update:model-value="handleCancel"
      @close="handleClose"
    >
      <ElForm ref="formRef" :model="form" :rules="rules" label-width="auto">
        <ElFormItem prop="images">
            <ElSpace wrap>
            <div v-for="(item,index) in form.images" 
                class="w-[100px] h-[100px] relative">
                <ElImage
                    class="rounded-lg"
                    :src="item"
                    :preview-src-list="[item]"
                    >
                </ElImage>
                <div @click="deleteImages(index)"  class="absolute top-0 right-0 cursor-pointer">
                    <ArtSvgIcon icon="solar:close-circle-bold" class="text-red-500 text-2xl"/>
                </div>
            </div>
            <ArtUpload
            @success="handleImages"
            >
                <div class="w-[100px] h-[100px] flex items-center justify-center border-1 border-solid rounded-lg cursor-pointer">
                <ArtSvgIcon icon="solar:upload-square-bold" class="text-2xl"/>
                </div>
            </ArtUpload>
            </ElSpace>
      </ElFormItem>
      </ElForm>
      <template #footer>
        <ElButton @click="handleCancel">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </template>
    </ElDialog>
  </template>
  
<script setup lang="ts">
import { fetchPostDistributeSettlement } from '@/api/distribute';
import type { FormInstance, FormRules } from 'element-plus'


interface Props {
    visible: boolean
    id?: number | null
}

interface Emits {
    (e: 'update:visible', value: boolean): void
    (e: 'submit'): void
}

const props = withDefaults(defineProps<Props>(), {
    visible: false,
    id: 0,
})

const emit = defineEmits<Emits>()

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
const form = reactive<Distribute.Params.Settlement>({
    id: 0, // 权限ID
    images: [],
})


const rules = reactive<FormRules>({
    images: [
        { required: true, message: '请选择上传报单文件', trigger: 'blur' },
    ],
})
const handleImages = (e:string) => {
  form.images.push(e)
}

const deleteImages = (i:number) => {
    form.images.splice(i,1)
}

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
    Object.assign(form, {
        id: props.id!,
    })
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
        await fetchPostDistributeSettlement(form)

        ElMessage.success('提交成功')
        emit('submit')
        handleClose()
        handleCancel()
    } catch (error) {
        console.log('表单验证失败:', error)
    }
}
</script>