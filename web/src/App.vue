<template>
  <div>
    <div id="overlay" onclick="overlayOff()"></div>
    <nav class="navbar navbar-expand-lg navbar-light bg-light static-top border-bottom box-shadow">
      <router-link to="/">
        <img src="https://content.parsableweb.com/parsableweb-logo-xsmall.png" alt="" width="100">
      </router-link>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarResponsive">
        <ul class="navbar-nav ml-auto">
          <li class="nav-item active">
            <router-link to="/" class="nav-link">Home</router-link>
          </li>
          <li class="nav-item">
            <router-link v-if="authenticated" to="/dashboard" class="nav-link">Latest Runs</router-link>
          </li>
          <li class="nav-item">
            <router-link v-if="authenticated" to="/jobs" class="nav-link">My Jobs</router-link>
          </li>
          <li class="nav-item">
            <router-link v-if="!authenticated" to="/login" class="btn btn-outline-primary">Sign In</router-link>
          </li>
          <!-- <li class="nav-item">
              <router-link v-if="authenticated" to="/faq" class="nav-link">FAQ</router-link>
          </li> -->
          <li class="nav-item">
              <router-link v-if="authenticated" to="/profile" class="nav-link">Profile</router-link>
          </li>
          <li class="nav-item">
              <router-link v-if="authenticated" to="/login" v-on:click.native="logout()" class="nav-link">Logout</router-link>
          </li>
          <!-- <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
            Account
            </a>
            <div class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
              <router-link v-if="authenticated" to="/profile" class="dropdown-item">Profile</router-link>
              <router-link v-if="authenticated" to="/login" v-on:click.native="logout()" class="dropdown-item">Logout</router-link>
            </div>
          </li> -->
        </ul>
        <small v-if="account.name">Hi, {{account.name}}</small>
      </div>
  </nav>
  <router-view @authenticated="setAuthenticated" :account="account"/>
  <footer class="container">
    <p class="float-right"><a href="#">Back to top</a></p>
    <p>
      &copy; 2021 braindumpnotes.com &middot; 
      <router-link to="/">Home</router-link> &middot; 
      <!-- <router-link to="/privacy">Privacy</router-link> &middot; 
      <router-link to="/terms">Terms of Use</router-link> &middot;  -->
      <router-link to="/contact">Contact</router-link></p>
  </footer>
  </div>
</template>

<script>
// import Vue from 'vue';
// import axios from 'axios';
// import _ from 'lodash';
import CommonMixin from './util/common';
// import AppConfig from './util/config';

export default {
    components: { },
    name: 'App',
    mixins: [CommonMixin],
    data() {
        return {
            account: {},
            authenticated: false,
            errors: []
        }
    },
    mounted() {
      // if(!this.authenticated) {
      //   if (_.includes(
      //     [
      //     "login",
      //     "register", 
      //     "disable",
      //     "contact",
      //     "get-started"
      //     ], this.$route.name)) 
      //   {
      //     this.setAccountData(null);
      //   } else {
      //     this.setAccountData("home");
      //   }
      // }
    },
    methods: {
        getAccount() {
          return this.account;
        },
        setAccountData(redirectOnError) {
          var that = this;
          this.backendGet("/account", redirectOnError).then(result => {
            that.account = result.data;
            that.authenticated = true;
          });
        },
        setAuthenticated(status) {
            this.authenticated = status;
            this.setAccountData();
        },
        logout() {
            this.deleteCookie("token");
            this.authenticated = false;
            this.account = {};
            this.$router.replace({ name: "home" });
        }
    }
}
</script>

<style>
/* #app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
} */

/* #nav {
  text-align: center;
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
} */
#logo {
  width: 400px;
  height: 200px;
  background-repeat: no-repeat;
  background-image: url('https://content.parsableweb.com/parsableweb-logo-small.png')
}
#logo-small {
  height: 80px;
  background-repeat: no-repeat;
  background-image: url('https://content.parsableweb.com/parsableweb-logo-xsmall.png')
}
footer {
  padding-top: 6rem;
  padding-bottom: 3rem;
}

footer p {
  margin-bottom: .25rem;
}
#overlay {
  /* position: absolute; */
  display: none;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0,0,0,0.5);
  z-index: 100;
}
</style>
