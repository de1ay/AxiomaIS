<template>
    <form id="panel_modal_franchises">
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.name"
          type="text"
          placeholder="Название">
        <icon name="building" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.official_name"
          type="text"
          placeholder="Официальное наименование">
        <icon name="building" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.full_official_name"
          type="text"
          placeholder="Полное официальное наименование">
        <icon name="building" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.license"
          type="text"
          placeholder="Лицензия">
        <icon name="id-card" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.city"
          type="text"
          placeholder="Город">
        <icon name="industry" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.legal_address"
          type="text"
          placeholder="Юридический адрес">
        <icon name="legal" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.inn"
          type="text"
          pattern="\[0-9]{10}"
          maxlength="10"
          placeholder="ИНН">
        <icon name="ticket" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.kpp"
          type="text"
          pattern="\[0-9]{9}"
          maxlength="9"
          placeholder="КПП">
        <icon name="sticky-note" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.bank_account"
          type="text"
          pattern="\[0-9]{20}"
          maxlength="20"
          placeholder="Р/C">
        <icon name="credit-card" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.bank"
          type="text"
          placeholder="Банк">
        <icon name="bank" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.correspondent_account"
          type="text"
          pattern="\[0-9]{19}"
          maxlength="19"
          placeholder="Кор. счёт">
        <icon name="money" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.bik"
          type="text"
          pattern="\[0-9]{9}"
          maxlength="9"
          placeholder="БИК">
        <icon name="exchange" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.ogrn"
          type="text"
          pattern="\[0-9]{13}"
          maxlength="13"
          placeholder="ОГРН">
        <icon name="tag" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.okved"
          type="text"
          placeholder="ОКВЭД">
        <icon name="info-circle" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.act"
          type="text"
          placeholder="Действует">
        <icon name="registered" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <input
          class="form-field__input"
          v-model.trim="editable_data.correspondent_address"
          type="text"
          placeholder="Адрес для корреспонденции">
        <icon name="map" scale="1.2"></icon>
      </div>
      <div class="form-field">
        <the-mask
          class="form-field__input"
          v-model.trim="editable_data.number"
          mask="+7 (###) ###-##-##"
          type="tel"
          placeholder="Телефон"></the-mask>
        <icon name="phone" scale="1.2"></icon>
      </div>
        <div class="form-field__actions">
          <button class="form-field__edit" @click.prevent="editFranchiseConfirm">Изменить</button>
          <button class="form-field__delete" @click.prevent="deleteFranchiseConfirm">Удалить</button>
        </div>
    </form>
</template>

<script>
  import 'vue-awesome/icons/building'
  import {TheMask} from 'vue-the-mask'
  export default {
    name: 'panelModalFranchisesView',
    props: ['franchises', 'staff', 'clinics', 'additional_data'],
    data () {
      return {
        editable_data: Object.assign({}, this.additional_data)
      }
    },
    methods: {
      editFranchiseConfirm () {
        this.$snotify.confirm('Изменить франчайзи?', 'Изменение', {
          timeout: 2000,
          showProgressBar: true,
          closeOnClick: false,
          pauseOnHover: true,
          buttons: [
            {text: 'Да', action: (notifyId) => { this.editFranchise(); this.$snotify.remove(notifyId) }},
            {text: 'Нет', action: (notifyId) => this.$snotify.remove(notifyId)}
          ]
        })
      },
      editFranchise () {
        this.$emit('editFranchise', this.editable_data, this.additional_data)
      },
      deleteFranchiseConfirm () {
        this.$snotify.confirm('Удалить франчайзи?', 'Удаление', {
          timeout: 2000,
          showProgressBar: true,
          closeOnClick: false,
          pauseOnHover: true,
          buttons: [
            {text: 'Да', action: (notifyId) => { this.$emit('deleteFranchise', this.additional_data); this.$snotify.remove(notifyId) }},
            {text: 'Нет', action: (notifyId) => this.$snotify.remove(notifyId)}
          ]
        })
      }
    },
    components: {TheMask}
  }
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>

<style scoped>

    #panel_modal_franchises {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    .form-field__delete {
        width: 45%;
        height: 38px;
        border-radius: 5px;
        border: none;
        background-color: #e74c3c;
        text-transform: uppercase;
        color: #fff;
        font-size: 18px;
        outline: none;
        cursor: pointer;
        transition: all 0.3s ease-in-out;
    }

    .form-field__delete:hover { background: #c0392b; }

    .form-field__edit {
        width: 45%;
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

    .form-field__edit:hover {  background: #2980b9; }

    .form-field__actions {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        margin: 15px 0 30px 0;
        width: 350px;
    }

</style>
