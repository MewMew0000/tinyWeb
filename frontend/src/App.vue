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
      </el-main>
  </el-container>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const activeIndex = ref('home')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>