<template>
  <el-container>
      <el-header>
        <el-menu
          :default-active="activeIndex"
          class="el-menu-demo"
          mode="horizontal"
          :ellipsis="false"
          @select="handleSelect"
        >
          <el-menu-item index="home">Home</el-menu-item>
          <el-menu-item index="currencyExchange">CurrencyExchange</el-menu-item>
          <el-menu-item index="news">News</el-menu-item>
          <el-menu-item index="login">Login</el-menu-item>
          <el-menu-item index="register">Register</el-menu-item>
          <el-menu-item index="logout">Logout</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
  </el-container>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();
const route = useRoute();
const activeIndex = ref(route.name?.toString() || 'home');

watch(route, (to) => {
  activeIndex.value = to.name?.toString() || 'home';
})

const handleSelect = (key: string) => {
  if (key == 'logout') {
    router.push({ name: 'home' });
  } else {
    router.push({ name: key.charAt(0).toUpperCase() + key.slice(1) });
  }
}
</script>