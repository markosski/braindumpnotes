
import axios from 'axios';
// import _ from 'lodash';
import AppConfig from '../util/config';

var CommonMixin = {
  data() {
    return {
      errors: [],
      notices: []
    }
  },
  methods:{
    getCookie(cname) {
      var name = cname + "=";
      var ca = document.cookie.split(';');

      for(var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
          c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
          return c.substring(name.length, c.length);
        }
      }
      return "";
    },
    deleteCookie(cname) {
      document.cookie = cname + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
    },
    onError(msg) {
      this.errors.push(msg);
    },
    onNotice(msg) {
      this.notices.push(msg);
    },
    backendGet(path) {
      var config = {
        // headers: { "Authorization": "bearer " + this.getCookie("token") }
      };
      return axios.get(AppConfig.BACKEND_ENDPOINT + path, config).then(result => {
        return result.data;
      }).catch(e => {
        var error = 'Unknown Error. Please try again later.';
        if (e.response.status === 403) {
          error = "User not authenticated"
        } else if (e.response.status === 503) {
          error = "Service is temporarily unavailable. Please try again later."
        }
        throw(error);
      });
    },
    backendDelete(path) {
      var config = {
        // headers: { "Authorization": "bearer " + this.getCookie("token") }
      };
      return axios.delete(AppConfig.BACKEND_ENDPOINT + path, config).catch(e => {
        var error = 'Unknown Error. Please try again later.';
        if (e.response.status === 403) {
          error = "User not authenticated"
        } else if (e.response.status === 503) {
          error = "Service is temporarily unavailable. Please try again later."
        }
        throw(error);
      });
    },
    backendPatch(path, data) {
      var config = {
        // headers: { "Authorization": "bearer " + this.getCookie("token") }
      };
      return axios.patch(AppConfig.BACKEND_ENDPOINT + path, data, config).catch(e => {
        var error = 'Unknown Error. Please try again later.';
        if (e.response.status === 403) {
          error = "User not authenticated"
        } else if (e.response.status === 503) {
          error = "Service is temporarily unavailable. Please try again later."
        }
        throw(error);
      });
    },
    backendPost(path, data) {
      var config = {
        // headers: { "Authorization": "bearer " + this.getCookie("token") }
      };
      return axios.post(AppConfig.BACKEND_ENDPOINT + path, data, config).catch(e => {
        var error = 'Unknown Error. Please try again later.';
        if (e.response.status === 403) {
          error = "User not authenticated"
        } else if (e.response.status === 503) {
          error = "Service is temporarily unavailable. Please try again later."
        }
        throw(error);
      });
    },
    overlayOn() {
      document.getElementById("overlay").style.display = "flex";
    },
    overlayOff() {
      document.getElementById("overlay").style.display = "none";
    }
  }
};
export default CommonMixin;