<template>
  <main role="main" class="container text-center">
    <h1>Account Profile</h1>
    <div class="row">
      <div v-if="errors && errors.length">
        <div v-bind:key="error" v-for="error of errors" class="alert alert-warning alert-dismissible fade show" role="alert">
            {{error}}
        </div>
      </div>
    </div>
    <div class="row">
      <div v-if="notices && notices.length">
        <div v-bind:key="notice" v-for="notice of notices" class="alert alert-success alert-dismissible fade show" role="alert">
            {{notice}}
        </div>
      </div>
    </div>
    <div class="row my-md-3 pb-md-5">
      <table class="table">
        <tr>
          <td>Name</td>
          <td>{{account.name}}</td>
        </tr>
        <tr>
          <td>Email</td>
          <td>{{account.email}}</td>
        </tr>
        <tr>
          <td>Current Plan</td>
          <td>{{account.plan.name}}</td>
        </tr>
      </table>
    </div>
  </main>
</template>

<script>
import axios from 'axios';
import _ from 'lodash';

export default {
    name: 'Profile',
    data() {
      return {
        account: {},
        foo: '',
        notices: [],
        errors: []
      }
    },
    methods: {
    },
    mounted() {
      this.account = this.$parent.$data.account;
    }
}
</script>