import json

feature_list=['wind_speed','weather_temperature_celsius','weather_relative_humidity','global_horizontal_radiation','diffuse_horizontal_radiation','wind_direction','weather_daily_rainfall']

'''
with open('type.txt', 'w') as outf:
    for fea in feature_list:
        outf.write(fea + '\t' + 'numeric\n')


values = [[0.325604473, 28.29973586],[-5.713217577, 44.62394078],[3.15324092, 101.870402],[1.158924697, 1539.433431],[0.342084, 712.4493],[13.30446847, 330.7176514],[0, 31.00001]]
with open('value.txt', 'w') as outf:
    for f,v in zip(feature_list, values):
        outf.write(f + '\tnumeric\t' + json.dumps({'min':v[0],'max':v[1]}) + '\n')
    outf.write('target\tnumeric\t' + json.dumps({'min':-0.00013,'max':6.06767}) + '\n')
'''

with open('features', 'w') as outf:
    md = {'features': [], 'target': 'Active_Power'}
    for fea in feature_list:
        md['features'].append({'name': fea, 'type': 'string'})
    outf.write(json.dumps(md)+'\n')
