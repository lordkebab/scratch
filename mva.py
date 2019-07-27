# playing with moving averages

import numpy as np
import talib

daily = np.random.rand(1, 280)[0]	# ndarray of 280 closing prices

output = talib.MA(daily, 30)	# 30-day moving avg
output2 = talib.MA(daily[60:], 30)	# 30-day moving avg starting at day 60

output = list(map(lambda x: round(x, 2), output))
output2 = list(map(lambda x: round(x, 2), output2))

# will they be the same?
for i in list(zip(output[60:], output2)):
    print(i)
	
# try it with EMA
ema_output = talib.EMA(daily, 30)
ema_output2 = talib.EMA(daily[60:], 30)

ema_output = list(map(lambda x: round(x, 2), ema_output))
ema_output2 = list(map(lambda x: round(x, 2), ema_output2))

for i in list(zip(ema_output[60:], ema_output2)):
    print(i)