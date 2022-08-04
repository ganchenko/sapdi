from io import StringIO
import csv
import pandas as pd
import json
def on_input(inData):
    data = StringIO(inData.body) 
    var = json.dumps(inData.attributes) 
    result = json.loads(var)
    ABAP = result['ABAP']
    Fields = ABAP['Fields']
    columns = [] 
    for item in Fields:
        columns.append(item['Name'])   
    df = pd.read_csv(data, index_col = False, names = columns) 
    df_csv = df.to_csv(index = False, header = True)
    api.send('outString', df_csv)
api.set_port_callback('inData', on_input)