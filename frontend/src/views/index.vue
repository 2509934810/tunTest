<template>
  <div>
    <topbar/>
    <el-container>
      <el-aside>
        
      </el-aside>
      <el-main>
        <div id="caseShow" :style="{height: '300px', width: '900px'}"/>
      </el-main>
    </el-container>
  </div>
</template>
<script>
import topbar from '@/components/TopBar.vue'
export default {
  name: 'index',
  data() {
    return {
      name: ''
    }
  },
  components: {
    topbar
  },
  mounted(){
    this.drowLine()
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
    drowLine(){
      let show = this.$echarts.init(document.getElementById("caseShow"))
      show.setOption({
        title: { text: '在Vue中使用echarts' },
        tooltip: {},
        xAxis: {
            data: ["衬衫","羊毛衫","雪纺衫","裤子","高跟鞋","袜子"]
        },
        yAxis: {},
        series: [{
            name: '销量',
            type: 'bar',
            data: [5, 20, 36, 10, 10, 20]
        }]
      });
    }
  }
}
</script>