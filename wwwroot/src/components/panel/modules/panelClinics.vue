<template>
  <div id="panel_franchises">
    <div class="franchises">
      <vue-good-table
        :columns="clinics_columns"
        :rows="clinics"
        :paginate="true"
        :lineNumbers="true"
        :rowsPerPageText="rowsPerPage"
        :nextText="nextText"
        :prevText="prevText"
        :onClick="showBilling"
        :ofText="ofText">
        <div slot="emptystate">Клиники отсутствуют</div>
      </vue-good-table>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'panelClinics',
    props: ['clinics'],
    data () {
      return {
        ofText: 'из',
        nextText: 'След.',
        prevText: 'Пред.',
        rowsPerPage: 'Показывать строк:',
        clinics_columns: [
          {
            label: 'Название',
            field: 'name'
          },
          {
            label: 'Франчайзи',
            field: 'franchise.name'
          }
        ]
      }
    },
    methods: {
      showBilling (rowObj, index) {
        this.$emit('showModal', 'panel_clinics_view', rowObj)
      }
    }
  }
</script>

<style>

  #panel_franchises { margin:50px; }

  .franchises .table-bordered th, .franchises .table.table-bordered td {
    border-top: none !important;
    border-right: none !important;
    border-left: none !important;
    border-bottom: 0.5px rgba(52, 152, 219, 0.4) solid !important;
  }

  .franchises .table.condensed thead th, .franchises .table thead th {
    background-color: #3498db !important;
    color: #fff;
  }

  .franchises .good-table .table-footer {
    background: transparent;
    border: none;
  }

  .franchises .good-table .table-footer .pagination-controls {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
  }

  .franchises .table .sorting-asc, .franchises .table .sorting-desc {
    color: #fff !important;
    user-select: none;
  }

  .franchises .table .sorting-asc:after, .franchises .table .sorting-desc:after {
    border-top-color: #fff !important;
    border-bottom-color: #fff !important;
  }

  .franchises .good-table .table tbody th.line-numbers {
    background-color: transparent !important;
  }

  .franchises .good-table .table-footer .pagination-controls .info { font-size: 14px; }

</style>
