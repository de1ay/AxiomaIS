<template>
    <form id="panel_modal_clinics">
        <div class="form-field">
          <input
            class="form-field__input"
            v-model.trim="editable_data.name"
            type="text"
            placeholder="Название">
          <icon name="plus-circle" scale="1.2" class="form-field__icon--tv"></icon>
        </div>
        <div class="form-field">
          <multiselect
            class="form-field__input"
            v-model="editable_data.parent_info"
            :options="franchises"
            label="name"
            track-by="id"
            :allow-empty="false"
            :show-labels="false"
            placeholder="Франчайз"></multiselect>
          <icon name="building" scale="1.2" class="form-field__icon--handshake"></icon>
        </div>
        <div class="form-field__actions">
          <button class="form-field__edit" @click.prevent="editClinicConfirm">Изменить</button>
          <button class="form-field__delete" @click.prevent="deleteClinicConfirm">Удалить</button>
        </div>
    </form>
</template>

<script>
  import 'vue-awesome/icons/globe'
  import 'vue-awesome/icons/television'
  import 'vue-awesome/icons/bullhorn'
  export default {
    name: 'panelModalClinicsView',
    props: ['franchises', 'staff', 'clinics', 'additional_data'],
    data () {
      return {
        editable_data: Object.assign({}, this.additional_data)
      }
    },
    created () {
      this.editable_data.parent_info = this.franchises.filter(franchise => {
        return franchise.id === this.editable_data.parent
      })[0]
    },
    methods: {
      editClinicConfirm () {
        this.$snotify.confirm('Изменить клинику?', 'Изменение', {
          timeout: 2000,
          showProgressBar: true,
          closeOnClick: false,
          pauseOnHover: true,
          buttons: [
            {text: 'Да', action: (notifyId) => { this.editClinic(); this.$snotify.remove(notifyId) }},
            {text: 'Нет', action: (notifyId) => this.$snotify.remove(notifyId)}
          ]
        })
      },
      editClinic () {
        this.editable_data.parent = this.editable_data.parent_info.id
        this.$emit('editClinic', this.editable_data, this.additional_data)
      },
      deleteClinicConfirm () {
        this.$snotify.confirm('Удалить клинику?', 'Удаление', {
          timeout: 2000,
          showProgressBar: true,
          closeOnClick: false,
          pauseOnHover: true,
          buttons: [
            {text: 'Да', action: (notifyId) => { this.$emit('deleteClinic', this.additional_data); this.$snotify.remove(notifyId) }},
            {text: 'Нет', action: (notifyId) => this.$snotify.remove(notifyId)}
          ]
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

    .form-field__delete:hover {  background: #c0392b; }

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
        width: 300px;
    }

</style>
