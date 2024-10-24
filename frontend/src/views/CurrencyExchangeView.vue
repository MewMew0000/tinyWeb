<template>
    <el-container>
        <el-form :model="form" class="exchange-form">
            <el-form-item label="FromCurrency" label-width="80px">
                <el-select v-model="form.fromCurrency" placeholder="Choose Currency">
                    <el-option v-for="currency in currencies" :key="currency" :label="currency" :value="currency"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="ToCurrency" label-width="80px">
                <el-select v-model="form.toCurrency" placeholder="Choose Currency">
                    <el-option v-for="currency in currencies" :key="currency" :label="currency" :value="currency"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="Amount" label-width="80px">
                <el-input v-model="form.amount" placeholder="Input Amount"></el-input>
            </el-form-item>
            <el-form-item>  
                <el-button type="primary" @click="exchange">Exchange</el-button>  
            </el-form-item> 
        </el-form>
        <div v-if="result" class="result">  
            <p>Exchange result: {{ result }}</p>  
        </div>
    </el-container>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import axios from '../axios';

interface ExchangeRate {
    fromCurrency: string;
    toCurrency: string;
    rate: number;
}

const form = ref({
    fromCurrency: '',
    toCurrency: '',
    amount: 0,
});

const result = ref<number | null>(null);
const currencies = ref<string[]>([]);
const rates = ref<ExchangeRate[]>([]);

const fetchCurrencies = async () => {
    try {
        const response = await axios.get<ExchangeRate[]>('/exchangeRates');
        rates.value = response.data;
        currencies.value = [...new Set(response.data.map((rate: ExchangeRate) => [rate.fromCurrency, rate.toCurrency]).flat())];  
    }
    catch (error) {
        console.log('failed to fetch currencies',error);
    }
};

const exchange = () => {
    const rate = rates.value.find(
      (rate) => rate.fromCurrency === form.value.fromCurrency && rate.toCurrency === form.value.toCurrency
    )?.rate;
  
    if (rate) {
      result.value = form.value.amount * rate;
    } else {
      result.value = null;
    }
};

onMounted(fetchCurrencies);

</script>