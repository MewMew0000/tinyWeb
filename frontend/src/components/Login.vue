<template>
    <div class="auth-container">
        <el-form :model="form" class="auth-form" @submit.prevent="login">
            <el-form-item label="username" label-wideth="80px">
                <el-input v-model="form.username" placeholder="please input username" /> 
            </el-form-item>
            <el-form-item label="password" label-width="80px">  
                <el-input v-model="form.password" type="password" placeholder="please input password" />  
            </el-form-item>
                <el-button type="primary" native-type="submit">login</el-button> 
        </el-form>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '../store/auth';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';

const authStore = useAuthStore();
const router = useRouter();

const form = ref({
    username: '',
    password: ''
})

const login = async () => {
    try {
        await authStore.login(form.value.username, form.value.password);
        router.push({ name: 'News' });
    }
    catch {
        ElMessage.error("login failed please check username and password!");
    }
}
</script>