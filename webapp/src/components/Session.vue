<template>
  <div class="login-panel">
    <el-row class="login-tab-group">
      <el-col :span="12">
        <div class="text-center">
          <router-link class="none-decoration-a" :to="{name: 'login'}">Login</router-link>
        </div>
      </el-col>
      <el-col :span="12">
        <div class="text-center">
          <router-link class="none-decoration-a" :to="{name: 'register'}">Register</router-link>
        </div>
      </el-col>
    </el-row>
    <hr>
    <transition appear name="fade" mode="out-in">
      <router-view></router-view>
    </transition>
  </div>
</template>

<script>
  import Vue from 'vue'
  import {
    Row,
    Col
  } from 'element-ui'
  import 'element-ui/lib/theme-default/index.css'
  import 'bootstrap/dist/css/bootstrap.css'

  Vue.use(Row)
  Vue.use(Col)

  export default {
    name: 'session'
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  html body {
    height: 100%;
    width: 100%;
  }

  .login-panel {
    width: 300px;
    margin: auto;
    margin-top: 200px;
  }

  .none-decoration-a {
    text-decoration: none !important;
  }

  hr {
    margin-top: 10px;
    margin-bottom: 0px;
    clear: both;
    border: 1px;
    height: 1px;
    background-image: -webkit-linear-gradient(left, rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.15), rgba(0, 0, 0, 0));
    background-image: -moz-linear-gradient(left, rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.15), rgba(0, 0, 0, 0));
    background-image: -ms-linear-gradient(left, rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.15), rgba(0, 0, 0, 0));
    background-image: -o-linear-gradient(left, rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.15), rgba(0, 0, 0, 0));
  }

  .login-tab-group {
    width: 100%;
    height: 40px;
    font-size: x-large;
    font-weight: bolder;
  }

  a {
    color: #777;

  }
  .router-link-active {
    color: #1d90e6;
  }

  .fade-enter-active, .fade-leave-active {
    transition: opacity .3s
  }
  .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
    opacity: 0
  }
</style>
