<!-- 个人中心页面 -->
<template>
  <div class="w-full h-full p-0 bg-transparent border-none shadow-none">
    <div class="relative flex-b mt-2.5 max-md:block max-md:mt-1">
      <div class="w-112 mr-5 max-md:w-full max-md:mr-0">
        <div class="art-card-sm relative p-9 pb-6 overflow-hidden text-center">
          <img class="absolute top-0 left-0 w-full h-50 object-cover" src="@imgs/user/bg.webp" />
          <img
            class="relative z-10 w-20 h-20 mt-30 mx-auto object-cover border-2 border-white rounded-full"
            :src="userInfo.avatar!"
          />
          <ArtUpload @success="handleAvatar">
            <ArtSvgIcon icon="solar:camera-minimalistic-bold" 
            class="text-g-700 text-3xl relative z-10 left-[60px] top-[-30px]" />
          </ArtUpload>
          <h2 class="mt-5 text-xl font-normal">{{ userInfo.name }}</h2>
          <p class="mt-5 text-sm">{{ userInfo.description }}</p>

          <div class="w-75 mx-auto mt-7.5 text-left">
            <div class="mt-2.5">
              <ArtSvgIcon icon="solar:gamepad-bold" class="text-g-700 mr-2" />
              <ElTag type="primary" size="small">{{ userInfo.game }}</ElTag>
            </div>
            <div class="mt-2.5">
              <ArtSvgIcon icon="solar:bookmark-bold" class="text-g-700 mr-2" />
              <ElTag type="primary" size="small">{{ userInfo.title }}</ElTag>
            </div>
            <div class="mt-2.5">
              <ArtSvgIcon icon="ri:map-pin-line" class="text-g-700" />
              <span class="ml-2 text-sm">{{ address }}</span>
            </div>
            <div class="mt-2.5">
              <ArtSvgIcon icon="ri:dribbble-fill" class="text-g-700" />
              <span class="ml-2 text-sm">{{ userInfo.loginIp }}</span>
            </div>
          </div>
        </div>
      </div>
      <div class="flex-1 overflow-hidden max-md:w-full max-md:mt-3.5">
        <div class="art-card-sm">
          <h1 class="p-4 text-xl font-normal border-b border-g-300">基本设置</h1>

          <ElForm
            :model="form"
            class="box-border p-5 [&>.el-row_.el-form-item]:w-[calc(50%-10px)] [&>.el-row_.el-input]:w-full [&>.el-row_.el-select]:w-full"
            ref="ruleFormRef"
            :rules="rules"
            label-width="86px"
            label-position="top"
          >
            <ElRow>
              <ElFormItem label="昵称" prop="name">
                <ElInput v-model="form.name" :disabled="!isEdit" />
              </ElFormItem>
              <ElFormItem label="性别" prop="sex" class="ml-5">
                <ElSelect v-model="form.sex" :disabled="!isEdit">
                  <ElOption label="男" :value="SexType.Male" />
                  <ElOption label="女" :value="SexType.Female" />
                  <ElOption label="其他" :value="SexType.Other" />
                </ElSelect>
              </ElFormItem>
            </ElRow>

            <ElRow>
              <ElFormItem label="所在地" prop="address">
                <ElCascader :disabled="!isEdit"
                  style="width: 100%;"
                  v-model="form.address"
                  :options="region"
                  placeholder="请选择省市区"
                />
              </ElFormItem>
              <ElFormItem label="出生日期" prop="birthday" class="ml-5">
                <ElDatePicker :disabled="!isEdit"
                  style="width: 100%;"
                  v-model="form.birthday"
                  type="date"
                  placeholder="请选择出生日期"
                  format="YYYY/MM/DD"
                  value-format="x"
                />
              </ElFormItem>
            </ElRow>

            <ElFormItem label="相册" prop="tags">
              <ElSpace wrap>
                <div v-for="(item,index) in form.album" 
                    class="w-[100px] h-[100px] relative">
                    <ElImage
                      class="rounded-lg"
                      :src="item"
                      :preview-src-list="[item]"
                      >
                    </ElImage>
                    <div v-if="isEdit" @click="deleteAlbum(index)"  class="absolute top-0 right-0 cursor-pointer">
                      <ArtSvgIcon icon="solar:close-circle-bold" class="text-red-500 text-2xl"/>
                    </div>
                </div>
                <ArtUpload v-if="isEdit"
                @success="handleAlbum"
                >
                  <div class="w-[100px] h-[100px] flex items-center justify-center border-1 border-solid rounded-lg cursor-pointer">
                    <ArtSvgIcon icon="solar:upload-square-bold" class="text-2xl"/>
                  </div>
                </ArtUpload>
              </ElSpace>
            </ElFormItem>

            <ElFormItem label="个人介绍" prop="des" class="h-32">
              <ElInput type="textarea" :rows="4" v-model="form.description" :disabled="!isEdit" />
            </ElFormItem>

            <div class="flex-c justify-end [&_.el-button]:!w-27.5">
              <ElButton type="primary" class="w-22.5" v-ripple @click="edit">
                {{ isEdit ? '保存' : '编辑' }}
              </ElButton>
            </div>
          </ElForm>
        </div>
        <div class="art-card-sm my-5">
          <h1 class="p-4 text-xl font-normal border-b border-g-300">更改密码</h1>

          <ElForm :model="pwdForm" class="box-border p-5" label-width="86px" label-position="top">
            <ElFormItem label="当前密码" prop="oldPass">
              <ElInput
               placeholder="请输入当前密码"
                v-model="pwdForm.oldPass"
                type="password"
                :disabled="!isEditPwd"
                show-password
              />
            </ElFormItem>

            <ElFormItem label="新密码" prop="newPass">
              <ElInput
               placeholder="请输入新密码"
                v-model="pwdForm.newPass"
                type="password"
                :disabled="!isEditPwd"
                show-password
              />
            </ElFormItem>

            <ElFormItem label="确认新密码" prop="confirmPass">
              <ElInput
                placeholder="请输入确认新密码"
                v-model="pwdForm.confirmPass"
                type="password"
                :disabled="!isEditPwd"
                show-password
              />
            </ElFormItem>

            <div class="flex-c justify-end [&_.el-button]:!w-27.5">
              <ElButton type="primary" class="w-22.5" v-ripple @click="editPwd">
                {{ isEditPwd ? '保存' : '编辑' }}
              </ElButton>
            </div>
          </ElForm>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { SexType } from '@/enums/typeEnum';
import { useUserStore } from '@/store/modules/user'
import type { CascaderOption, FormInstance, FormRules } from 'element-plus'
import { provinceAndCityData,codeToText } from 'element-china-area-data';
import { fetchPostAccountChangePass, fetchPostAccountEdit } from '@/api/account';
defineOptions({ name: 'UserCenter' })
const region = provinceAndCityData as CascaderOption[]


const userStore = useUserStore()
const userInfo = computed(() => userStore.getUserInfo)
const address = computed(()=>{
    if (userInfo.value.address!.length  == 2) {
        return codeToText[userInfo.value.address![0]]  + '/' + codeToText[userInfo.value.address![1]]
    }
    return codeToText[userInfo.value.address![0]]
})

const isEdit = ref(false)
const isEditPwd = ref(false)
const date = ref('')
const ruleFormRef = ref<FormInstance>()

/**
 * 用户信息表单
 */
const form = reactive<Account.Params.Model>({
  avatar: userInfo.value.avatar!,
  name: userInfo.value.name!,
  address: userInfo.value.address!,
  sex: userInfo.value.sex!,
  description: userInfo.value.description!,
  birthday: userInfo.value.birthday!,
  album: userInfo.value.album!,
})

/**
 * 密码修改表单
 */
const pwdForm = reactive<Account.Params.ChangePass>({
  oldPass: null,
  newPass: null,
  confirmPass: null,
})

/**
 * 表单验证规则
 */
const rules = reactive<FormRules>({
})


onMounted(() => {
  getDate()
})

/**
 * 根据当前时间获取问候语
 */
const getDate = () => {
  const h = new Date().getHours()

  if (h >= 6 && h < 9) date.value = '早上好'
  else if (h >= 9 && h < 11) date.value = '上午好'
  else if (h >= 11 && h < 13) date.value = '中午好'
  else if (h >= 13 && h < 18) date.value = '下午好'
  else if (h >= 18 && h < 24) date.value = '晚上好'
  else date.value = '很晚了，早点睡'
}

const handleAvatar = async (e:string)=>{
  // TODO: 调用新增/编辑接口
  await fetchPostAccountEdit({
    avatar: e,
    name: userInfo.value.name!,
    address: userInfo.value.address!,
    sex: userInfo.value.sex!,
    description: userInfo.value.description!,
    birthday: userInfo.value.birthday!,
    album: userInfo.value.album!,
  })
    
  userStore.setUserInfo({
    id: userInfo.value.id!,
    name: userInfo.value.name!,
    address: userInfo.value.address!,
    sex:  userInfo.value.sex!,
    avatar: e,
    album: userInfo.value.album!,
    commission: userInfo.value.commission!,
    title: userInfo.value.title!,
    game: userInfo.value.game!,
    rate: userInfo.value.rate!,
    description: userInfo.value.description!,
    loginIp: userInfo.value.loginIp!,
    birthday: userInfo.value.birthday!,
  })
  ElMessage.success('更新头像成功')
}

const handleAlbum = (e:string) => {
  form.album.push(e)
}
const deleteAlbum = (i:number) => {
  form.album.splice(i,1)
}

/**
 * 切换用户信息编辑状态
 */
const edit = async () => {
  
  if (isEdit.value) {
    try {
      // TODO: 调用新增/编辑接口
      await fetchPostAccountEdit(form)
      ElMessage.success('更新成功')
      userStore.setUserInfo({
        id: userInfo.value.id!,
        name: form.name!,
        address: form.address!,
        sex:  form.sex!,
        avatar: userInfo.value.avatar!,
        birthday: form.birthday!,
        album: form.album!,
        description: form.description!,

        commission: userInfo.value.commission!,
        title: userInfo.value.title!,
        game: userInfo.value.game!,
        rate: userInfo.value.rate!,
        loginIp: userInfo.value.loginIp!,
      })
    } catch (error) {
        console.log('表单验证失败:', error)
    }
  }
  isEdit.value = !isEdit.value
}

/**
 * 切换密码编辑状态
 */
const editPwd = async () => {
  if (isEditPwd.value) {
    try {
      await fetchPostAccountChangePass(pwdForm)
      ElMessage.success('更新成功')
    } catch (error) {
      
    }
  }
  
  isEditPwd.value = !isEditPwd.value
}
</script>
