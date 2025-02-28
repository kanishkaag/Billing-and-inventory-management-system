# %%
# Cell 1: Import necessary libraries
import pandas as pd
import matplotlib.pyplot as plt
from sklearn.preprocessing import StandardScaler, OneHotEncoder
from sklearn.compose import ColumnTransformer
from statsmodels.tsa.statespace.sarimax import SARIMAX


# %%
# Cell 2: Load the data
# Assuming the data is already loaded
df = pd.read_csv("1.csv")


# %%
# Cell 3: Feature engineering - Add year and week of the year columns
df['transaction_date'] = pd.to_datetime(df['transaction_date'])
df['week_of_year'] = df['transaction_date'].dt.isocalendar().week
df['year'] = df['transaction_date'].dt.year


# %%
# Cell 4: Function to aggregate sales data weekly
def aggregate_weekly_sales(df_product):
    weekly_sales = df_product.groupby(['year', 'week_of_year']).agg({'quantity': 'sum'}).reset_index()
    weekly_sales['date'] = pd.to_datetime(weekly_sales['year'].astype(str) + '-' + weekly_sales['week_of_year'].astype(str) + '-1', format='%Y-%W-%w')
    weekly_sales.set_index('date', inplace=True)
    
    # Exogenous variables: Here, we'll just use 'discount_applied' and 'product_stock' as features for simplicity.
    exogenous = df_product.groupby(['year', 'week_of_year']).agg({'discount_applied': 'mean', 'product_stock': 'mean'}).reset_index()
    exogenous['date'] = pd.to_datetime(exogenous['year'].astype(str) + '-' + exogenous['week_of_year'].astype(str) + '-1', format='%Y-%W-%w')
    exogenous.set_index('date', inplace=True)
    
    return weekly_sales['quantity'], exogenous[['discount_applied', 'product_stock']]


# %%
# Cell 5: Feature engineering and preprocessing
def feature_engineering_and_preprocessing(df_product):
    df_product['week_of_year'] = df_product['transaction_date'].dt.isocalendar().week
    df_product['year'] = df_product['transaction_date'].dt.year
    return df_product


# %%
import pickle
from statsmodels.tsa.statespace.sarimax import SARIMAX

def train_and_save_sarimax_model(df, product_id, order=(1, 1, 1), seasonal_order=(1, 1, 1, 52)):
    # Filter the data for the specific product_id
    df_product = df[df['product_id'] == product_id]

    if df_product.empty:
        raise ValueError(f"No data found for product ID {product_id}")

    # Feature engineering and preprocessing
    df_product = feature_engineering_and_preprocessing(df_product)

    # Aggregate the sales data to weekly sales
    weekly_sales, exogenous_weekly = aggregate_weekly_sales(df_product)

    # Check if we have data to train the model
    if len(weekly_sales) == 0 or len(exogenous_weekly) == 0:
        raise ValueError(f"No valid sales or exogenous data available for product ID {product_id}")

    # Train the SARIMAX model
    model = SARIMAX(weekly_sales, exog=exogenous_weekly, order=order, seasonal_order=seasonal_order)
    model_fit = model.fit()


    
    return model_fit


# %%
import matplotlib.pyplot as plt
import pandas as pd

# Function to plot the forecast
def plot_sales_forecast(model_fit, df, product_id, forecast_steps=12):
    df_product = df[df['product_id'] == product_id]
    df_product = feature_engineering_and_preprocessing(df_product)

    # Aggregate the sales data to weekly sales
    weekly_sales, exogenous_weekly = aggregate_weekly_sales(df_product)

    # Forecasting the next 12 weeks
    forecast = model_fit.forecast(steps=forecast_steps, exog=exogenous_weekly[-forecast_steps:])

    # Plotting the historical sales and the forecasted sales
    plt.figure(figsize=(10, 6))
    
    # Plot historical sales
    plt.plot(weekly_sales.index, weekly_sales, label='Historical Sales', marker='o')
    
    # Generate the forecasted dates
    forecast_dates = pd.date_range(weekly_sales.index[-1], periods=forecast_steps + 1, freq='W')[1:]
    
    # Plot forecasted sales
    plt.plot(forecast_dates, forecast, color='red', label='Forecasted Sales', marker='x')
    
    # Add titles and labels
    plt.title(f'Sales Forecast for Product ID {product_id} - Next {forecast_steps} Weeks')
    plt.xlabel('Date')
    plt.ylabel('Quantity Sold')
    plt.legend()
    plt.grid()
    plt.show()

    return forecast


# %%
# Example usage


# Train and save the model for a specific product
model_fit = train_and_save_sarimax_model(df, product_id=4175, order=(1, 1, 1), seasonal_order=(1, 1, 1, 52))


# %%
import sys
product_id = int(sys.argv[1])
plot_sales_forecast(model_fit, df, product_id)


