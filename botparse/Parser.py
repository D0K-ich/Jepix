import time
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
import pickle
from config import email, password

options = Options() # присваиваю переменной options модуль Options из селениума 
options.add_argument('--headless') #избавляюсь от графической оболочки
driver = webdriver.Chrome(options=options) #запускаю драйвер

def login(driver): #функция для авторизации на сайте
    driver.maximize_window()
    driver.get('https://app.unityspace.ru/login')
    email_input = driver.find_element(By.NAME, 'email')
    email_input.clear()
    email_input.send_keys(email)
    password_input = driver.find_element(By.NAME, 'password')
    password_input.clear()
    password_input.send_keys(password)
    password_input.send_keys(Keys.ENTER)
    driver.implicitly_wait(10)
    page = driver.find_element(By.PARTIAL_LINK_TEXT, 'Project JEPIX').click()
    driver.implicitly_wait(10)
    tasks = driver.find_element(By.PARTIAL_LINK_TEXT, 'Main product').click()
    time.sleep(1)


def Backend_tasks_Pars(driver): # Парсинг Backend задач
    login(driver) #Вызываю функцию login для авторизации
    Backend = driver.find_element(By.XPATH, '//*[@id="app"]/section/div/div/div/div/div[2]/div/div[1]/div/div[1]/div[1]/p')
    BackendTasks = driver.find_element(By.ID, 'tasks-45705')
    Backend_tasks_allerts = f'Задачи на {Backend.text}:' + '\n' + BackendTasks.text + '\rНе проебись с дедлайнами!'
    return Backend_tasks_allerts
    driver.close() #Закрываю и выключаю драйвер
    driver.quit()
    
def BotBarse_tasks_Pars(driver): # Парсинг BotParse задач
    login(driver)
    Botparse = driver.find_element(By.XPATH, '//*[@id="app"]/section/div/div/div/div/div[2]/div/div[2]/div/div[1]/div[1]/p')
    BotparseTasks = driver.find_element(By.ID, 'tasks-45707')
    Bot_and_Parse_tasks_allerts = f'Задачи на {Botparse.text}:' + '\n' + BotparseTasks.text + '\n' + 'Не проебись с дедлайнами(Вас особенно касается)!!'
    return Bot_and_Parse_tasks_allerts
    driver.close()
    driver.quit()
    
def Frontend_tasks_Pars(driver): # Парсинг Frontend задач
    login(driver)
    Fronted = driver.find_element(By.XPATH, '//*[@id="app"]/section/div/div/div/div/div[2]/div/div[3]/div/div[1]/div[1]/p')
    FrontedTasks = driver.find_element(By.ID, 'tasks-45706')
    Frontend_tasks_allerts = f'Задачи на {Fronted.text}:' + '\n' + FrontedTasks.text+ '\n' + 'Не проебись с дедлайнами!!'
    return Frontend_tasks_allerts
    driver.close()
    driver.quit()

def Done_tasks_Pars(driver): # Парсинг Выполненных задач
    login(driver)
    Done = driver.find_element(By.XPATH, '//*[@id="app"]/section/div/div/div/div/div[2]/div/div[4]/div/div[1]/div[1]/p')
    DoneTasks = driver.find_element(By.ID, 'tasks-45704')
    Done_tasks_allerts = f'{Done.text}:' + '\n' + DoneTasks.text
    return Done_tasks_allerts
    driver.close()
    driver.quit()

#Первым делом избавляемся от графической оболочки и запускаем сам драйвер
#Затем авторизовываемся, используя функцию login, в параметрах которой находится драйвер
#После авторизации выполняется сам парсинг, разделенный по модулям, соответственно каждому блоку с задачами

