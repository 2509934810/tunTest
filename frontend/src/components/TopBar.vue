<template>
  <div>
    <el-menu
      :default-active="activeIndex2"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b">
      <el-menu-item index="1">测试中心</el-menu-item>
      <el-submenu index="2">
        <template slot="title">测试处理</template>
        <el-menu-item index="2-1">开始测试</el-menu-item>
        <el-menu-item index="2-2">配置管理</el-menu-item>
        <el-menu-item index="2-3">退出登陆</el-menu-item>
      </el-submenu>
      <el-submenu index="3">
          <template slot="title">测试区域</template>
          <el-menu-item index="3-1">最近一次</el-menu-item>
          <el-menu-item index="3-2">最近十次</el-menu-item>
          <el-menu-item index="3-3">最近三十次</el-menu-item>
          <el-menu-item index="3-3">最近一百次</el-menu-item>
      </el-submenu>
      <el-menu-item index="4">输入</el-menu-item>
      <el-menu-item index="4"><a>订单管理</a></el-menu-item>
    </el-menu>
  </div>
</template>
<script>
export default {
  name: 'topbar',
  data() {
      return {
        activeIndex: '1',
        activeIndex2: '1'
      };
    },
    methods: {
      handleSelect(key, keyPath) {
        console.log(key, keyPath);
      }
    }
}
</script>