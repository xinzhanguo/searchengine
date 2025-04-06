<template>

  <el-row>
    <el-col :span="12">
      <el-card>
        <template #header>
          github
        </template>
        <el-link type="success" target="_blank" href="https://github.com/xinzhanguo/searchengine">
          https://github.com/xinzhanguo/searchengine
        </el-link>
      </el-card>
    </el-col>

  </el-row>


</template>

<script>
export default {
  name: 'document.vue',
}
</script>

<style scoped>
.el-col{
  padding:10px;
}
</style>
