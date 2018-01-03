<template>
  <form id="panel_modal_clinics">
    <div class="form-field">
        <input
        class="form-field__input"
        v-model.trim="name"
        type="text"
        placeholder="Название">
        <icon name="plus-circle" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <multiselect
        class="form-field__input"
        v-model="franchise"
        :options="franchises"
        label="name"
        track-by="id"
        :allow-empty="false"
        :show-labels="false"
        placeholder="Франчайзи"></multiselect>
      <icon name="building" scale="1.2"></icon>
      </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="city"
        type="text"
        placeholder="Город">
      <icon name="industry" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="address"
        type="text"
        placeholder="Адрес">
      <icon name="map" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <multiselect
        class="form-field__input"
        v-model="active_info"
        :options="billing_options"
        track-by="id"
        label="label"
        :allow-empty="false"
        :show-labels="false"
        placeholder="Может принимать оплату"></multiselect>
      <icon name="money" scale="1.2"></icon>
    </div>
    <div class="form-field__actions">
        <input class="form-field__submit" @click.prevent="addClinic" type="submit" value="Создать"/>
    </div>
  </form>
</template>

<script>
  import 'vue-awesome/icons/plus-circle'
  import 'vue-awesome/icons/building'
  import 'vue-awesome/icons/money'
  import 'vue-awesome/icons/map'
  import 'vue-awesome/icons/industry'
  export default {
    name: 'panelModalClinicAdd',
    props: ['franchises', 'clinics', 'staff'],
    data () {
      return {
        name: '',
        franchise: '',
        address: '',
        active_info: '',
        city: '',
        billing_options: [
          {'id': true, 'label': 'Да'},
          {'id': false, 'label': 'Нет'}
        ]
      }
    },
    methods: {
      addClinic () {
        this.$emit('addClinic', {
          name: this.name,
          address: this.address,
          is_active: this.active_info.id,
          city: this.city,
          franchise: this.franchise,
          franchise_id: this.franchise.id
        })
      }
    }
  }
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>

<style scoped>

    #panel_modal_clinics {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    .form-field__submit {
        width: 100%;
        height: 38px;
        border-radius: 5px;
        border: none;
        background-color: #3498db;
        text-transform: uppercase;
        color: #fff;
        font-size: 18px;
        outline: none;
        cursor: pointer;
        transition: all 0.3s ease-in-out;
    }

    .form-field__submit:hover {  background: #2980b9; }

    .form-field__actions {
        display: flex;
        flex-direction: row;
        justify-content: flex-start;
        align-items: center;
        margin: 15px 0 30px 0;
        width: 350px;
    }

</style>
