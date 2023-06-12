from parser import parse
import asyncio


async def analyze(city_id: int, req_type: str):
    metrics = await parse(city_id, req_type)
    if city_id == 1:
        keys = list(metrics.keys())
        size = len(metrics) // 4

        saryarka = {key: metrics[key] for key in keys[:size]}
        almata = {key: metrics[key] for key in keys[size:size * 2]}
        baikonur = {key: metrics[key] for key in keys[size * 2:size * 3]}
        esil = {key: metrics[key] for key in keys[size * 3:]}
        if req_type == "temperature":
            saryarka_temp, saryarka_hum = calculate(saryarka, req_type)
            almata_temp, almata_hum = calculate(almata, req_type)
            baikonur_temp, baikonur_hum = calculate(baikonur, req_type)
            esil_temp, esil_hum = calculate(esil, req_type)
            saryarka_dict = {
                'city': city_id,
                'district': 'saryarka',
                'temperature': saryarka_temp,
                'humidity': saryarka_hum
            }
            almata_dict = {
                'city': city_id,
                'district': 'almatynskiy',
                'temperature': almata_temp,
                'humidity': almata_hum
            }
            baikonur_dict = {
                'city': city_id,
                'district': 'baikonur',
                'temperature': baikonur_temp,
                'humidity': baikonur_hum
            }
            esil_dict = {
                'city': city_id,
                'district': 'esil',
                'temperature': esil_temp,
                'humidity': esil_hum
            }
        else:
            saryarka_co, saryarka_pm = calculate(saryarka, req_type)
            almata_co, almata_pm = calculate(almata, req_type)
            baikonur_co, baikonur_pm = calculate(baikonur, req_type)
            esil_co, esil_pm = calculate(esil, req_type)

            saryarka_dict = {
                'city': city_id,
                'district': 'saryarka',
                'co2': saryarka_co,
                'pm25': saryarka_pm
            }
            almata_dict = {
                'city': city_id,
                'district': 'almatynskiy',
                'co2': almata_co,
                'pm25': almata_pm
            }
            baikonur_dict = {
                'city': city_id,
                'district': 'baikonur',
                'co2': baikonur_co,
                'pm25': baikonur_pm
            }
            esil_dict = {
                'city': city_id,
                'district': 'esil',
                'co2': esil_co,
                'pm25':  esil_pm
            }

    elif city_id == 2:
        keys = list(metrics.keys())
        size = len(metrics) // 5

        ksht = {key: metrics[key] for key in keys[:size]}
        zashita = {key: metrics[key] for key in keys[size:size * 2]}
        ulbinsk = {key: metrics[key] for key in keys[size * 2:size * 3]}
        tsentr = {key: metrics[key] for key in keys[size * 3:]}
        zavodsk = {key: metrics[key] for key in keys[size * 4:]}
        if req_type == "temperature":
            ksht_temp, ksht_hum = calculate(ksht, req_type)
            zashita_temp, zashita_hum = calculate(zashita, req_type)
            ulbinsk_temp, ulbinsk_hum = calculate(ulbinsk, req_type)
            zavodsk_temp, zavodsk_hum = calculate(zavodsk, req_type)
            tsentr_temp, tsentr_hum = calculate(tsentr, req_type)
            ksht_dict = {
                'city': city_id,
                'district': 'ksht',
                'temperature': ksht_temp,
                'humidity': ksht_hum
            }
            zashita_dict = {
                'city': city_id,
                'district': 'zashita',
                'temperature': zashita_temp,
                'humidity': zashita_hum
            }
            ulbinsk_dict = {
                'city': city_id,
                'district': 'ulbinsk',
                'temperature': ulbinsk_temp,
                'humidity': ulbinsk_hum
            }
            zavodsk_dict = {
                'city': city_id,
                'district': 'zavodsk',
                'temperature': zavodsk_temp,
                'humidity': zavodsk_hum
            }
            tsentr_dict = {
                'city': city_id,
                'district': 'tsentr',
                'temperature': tsentr_temp,
                'humidity': tsentr_hum
            }
        else:
            ksht_co, ksht_pm = calculate(ksht, req_type)
            zashita_co, zashita_pm = calculate(zashita, req_type)
            ulbinsk_co, ulbinsk_pm = calculate(ulbinsk, req_type)
            zavodsk_co, zavodsk_pm = calculate(zavodsk, req_type)
            tsentr_co, tsentr_pm = calculate(tsentr, req_type)

            ksht_dict = {
                'city': city_id,
                'district': 'ksht',
                'co2': ksht_co,
                'pm25': ksht_pm
            }
            zashita_dict = {
                'city': city_id,
                'district': 'zashita',
                'co2': zashita_co,
                'pm25': zashita_pm
            }
            ulbinsk_dict = {
                'city': city_id,
                'district': 'ulbinsk',
                'co2': ulbinsk_co,
                'pm25': ulbinsk_pm
            }
            zavodsk_dict = {
                'city': city_id,
                'district': 'zavodsk',
                'co2': zavodsk_co,
                'pm25':zavodsk_pm
            }
            tsentr_dict = {
                'city': city_id,
                'district': 'tsentr',
                'co2': tsentr_co,
                'pm25':tsentr_pm
            }

    elif city_id == 3:
        keys = list(metrics.keys())
        size = len(metrics) // 2

        atr_zap = {key: metrics[key] for key in keys[:size]}
        atr_vos = {key: metrics[key] for key in keys[size:size * 2]}

        if req_type == "temperature":
            atr_zap_temp, atr_zap_hum = calculate(atr_zap, req_type)
            atr_vos_temp, atr_vos_hum = calculate(atr_vos, req_type)

            atr_zap_dict = {
                'city': city_id,
                'district': 'zapad',
                'temperature': atr_zap_hum,
                'humidity': atr_zap_hum
            }
            atr_vos_dict = {
                'city': city_id,
                'district': 'vostok',
                'temperature': atr_vos_temp,
                'humidity': atr_vos_hum
            }

        else:
            atr_zap_co, atr_zap_pm = calculate(atr_zap, req_type)
            atr_vos_co, atr_vos_pm = calculate(atr_vos, req_type)

            atr_zap_dict = {
                'city': city_id,
                'district': 'zapad',
                'co2': atr_zap_co,
                'pm25': atr_zap_pm
            }
            atr_vos_dict = {
                'city': city_id,
                'district': 'vostok',
                'co2': atr_vos_co,
                'pm25': atr_vos_pm
            }

    elif city_id == 4:
        keys = list(metrics.keys())
        size = len(metrics) // 2

        sem_sev = {key: metrics[key] for key in keys[:size]}
        sem_yug = {key: metrics[key] for key in keys[size:size * 2]}

        if req_type == "temperature":
            sem_yug_temp, sem_yug_hum = calculate(sem_yug, req_type)
            sem_sev_temp, sem_sev_hum = calculate(sem_sev, req_type)
            sem_sev_dict = {
                'city': city_id,
                'district': 'sever',
                'temperature': sem_sev_temp,
                'humidity': sem_sev_hum
            }
            sem_yug_dict = {
                'city': city_id,
                'district': 'yug',
                'temperature': sem_yug_temp,
                'humidity': sem_yug_hum
            }
        else:
            sem_yug_co, sem_yug_pm = calculate(sem_yug, req_type)
            sem_sev_co, sem_sev_pm = calculate(sem_sev, req_type)
            sem_sev_dict = {
                'city': city_id,
                'district': 'sever',
                'co2': sem_sev_co,
                'pm25': sem_sev_pm
            }
            sem_yug_dict = {
                'city': city_id,
                'district': 'yug',
                'co2': sem_yug_co,
                'pm25': sem_yug_co
            }


def calculate(dicti, req_type):
    size = len(dicti)
    for item in dicti.values():
        dicti_temperature = 0
        dicti_humidity = 0
        dicti_co = 0
        dicti_pm = 0
        if req_type == "temperature":
            dicti_temperature += item['temperature']
            dicti_humidity += item['humidity']
            mean_temp = dicti_temperature / size
            mean_hum = dicti_humidity / size
            return mean_temp, mean_hum
        else:
            dicti_co += item['co2']
            dicti_pm += item['pm25']
            mean_co = dicti_co / size
            mean_pm = dicti_pm / size
            return mean_co, mean_pm


async def main():
    await analyze(2, "temperature")


loop = asyncio.get_event_loop()
loop.run_until_complete(main())
