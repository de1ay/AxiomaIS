<template>
  <div class="SignIn">
    <form class="signIn__form">
      <div class="form-field">
        <input class="form-field__input" id="signIn__form-login" v-model.trim="login" type="text" placeholder="Логин">
        <icon name="user" scale="2" class="form-field__icon--user"></icon>
      </div>
      <div class="form-field">
        <input class="form-field__input" id="signIn__form-password" v-model.trim="password" type="password" placeholder="Пароль">
        <icon name="lock" scale="2" class="form-field__icon--lock"></icon>
      </div>
      <div class="form-field__actions">
        <input @click.prevent="authorize" class="form-field__submit" type="submit" value="Войти"/>
      </div>
    </form>
  </div>
</template>

<script>
  import axios from 'axios'
  import 'vue-awesome/icons/user'
  import 'vue-awesome/icons/lock'
  export default {
    name: 'ignIn',
    data () {
      return {
        login: '',
        password: ''
      }
    },
    created () {
      if (this.$cookie.get('sid')) {
        this.$router.push({name: 'panel_main'})
      }
    },
    methods: {
      authorize () {
        if (this.login.length < 3) {
          this.$snotify.error('Логин короче 3 символов', 'Ошибка!', {
            timeout: 2000,
            showProgressBar: true,
            closeOnClick: true,
            pauseOnHover: true
          })
          return
        }
        if (this.password.length < 8) {
          this.$snotify.error('Пароль короче 8 символов', 'Ошибка!', {
            timeout: 2000,
            showProgressBar: true,
            closeOnClick: true,
            pauseOnHover: true
          })
          return
        }
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.get('https://api.axiomais.ru/token.get?login=' + this.login + '&password=' + this.password).then(resp => {
              if (resp.data.code !== 200) {
                /*eslint-disable */
                reject({
                  title: 'Ошибка!',
                  body: resp.data.message,
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
                /*eslint-enable */
              }
              this.$cookie.set('sid', resp.data.message.token, {expires: 31})
              this.$router.push({name: 'panel_main'})
              resolve({
                title: 'Успешно',
                body: 'Авторизация успешна',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: resp.data.message,
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          })
        )
      }
    }
  }
</script>

<style scoped>

  .SignIn {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
    width: 100vw;
    background-image: url('../assets/images/axioma.jpg');
    background-size: cover;
  }

  .signIn__form {
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
    height: 304px;
    width: 598px;
    background-color: #ecf0f1;
    border-radius: 5px;
  }

  .SignIn .signIn__form .form-field {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    width: 534px;
    height: 55px;
    border-radius: 5px;
    background-color: #ffffff;
  }

  .form-field .fa-icon { color: #95a5a6; }

  .SignIn .signIn__form .form-field .form-field__icon--user {
    width: 21px;
    height: 26px;
  }

  .SignIn .signIn__form .form-field .form-field__icon--lock {
    margin-top: 4px;
    width: 21px;
    height: 26px;
  }

  .SignIn .signIn__form .form-field__input {
    width: 474px;
    height: 52px;
    font-size: 24px;
    color: #95a5a6;
    background-color: #fff;
    border: none;
    outline: none;
  }

  .form-field__input::-webkit-input-placeholder { color: #95a5a6; }
  .form-field__input:-moz-placeholder { color: #95a5a6; }
  .form-field__input::-moz-placeholder { color: #95a5a6; }
  .form-field__input:-ms-input-placeholder { color: #95a5a6; }

  .SignIn .signIn__form .form-field__submit {
    width: 100%;
    height: 55px;
    border-radius: 5px;
    border: none;
    background-color: #3498db;
    text-transform: uppercase;
    color: #fff;
    font-size: 30px;
    outline: none;
    cursor: pointer;
    transition: all 0.3s ease-in-out;
  }

  .SignIn .signIn__form .form-field__submit:hover {  background: #2980b9; }

  .SignIn .signIn__form .form-field__actions {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    width: 534px;
  }

</style>
