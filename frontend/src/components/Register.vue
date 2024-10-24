<template>
    <div class="auth-container">
      <el-form :model="form" class="auth-form" @submit.prevent="register">
        <el-form-item label="username" label-width="80px">
            <el-input v-model="form.username" placeholder="please input username" />
        </el-form-item>
        <el-form-item label="password" label-width="80px">
            <el-input v-model="form.password" type="password" placeholder="please input password" />
        </el-form-item>
            <el-button type="primary" native-type="submit">register</el-button>
      </el-form>
    </div>
    </template>
    
    <script setup lang="ts">
    import { ref } from 'vue';
    import { useRouter } from 'vue-router';
    import { useAuthStore } from '../store/auth';
    import { ElMessage } from 'element-plus';
    
    const form = ref({
      username: '',
      password: '',
    });
    
    const authStore = useAuthStore();
    const router = useRouter();
    
    const register = async () => {
      try {
        await authStore.register(form.value.username, form.value.password);
        router.push({ name: 'News' });
      } catch {
        ElMessage.error('register failed');
      }
    };
    </script>
    