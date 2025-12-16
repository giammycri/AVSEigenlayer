import warnings
warnings.filterwarnings("ignore")

import os

# Numero di script da generare
num_scripts = 5

# Contenuto base dello script
script_content = """
import syft as sy
from keras.datasets import mnist
from syft.service.user.user import ServiceRole
from syft.service.user.user import UserCreate

domain_client = sy.login(port="3331", email="info@openmined.org", password="changethis")
domain_client.users.create(UserCreate(email="alice{index}@openmined.org", name="Alice{index}", password="pw", role=ServiceRole.DATA_OWNER))
alice_do_{index} = domain_client.login(email="alice{index}@openmined.org", password="pw")

(train_X_a_{index}, train_y_a_{index}), (test_X_a_{index}, test_y_a_{index}) = mnist.load_data()

dataset = sy.Dataset(name="MNIST_{index}")

trainx = sy.Asset(name="Asset di Train")
trainx.set_obj(train_X_a_{index})
trainx.no_mock()

trainy = sy.Asset(name="Asset di etichette di Train")
trainy.set_obj(train_y_a_{index})
trainy.no_mock()

testx = sy.Asset(name="Asset di Test")
testx.set_obj(test_X_a_{index})
testx.no_mock()

testy = sy.Asset(name="Asset di eitchette di Test")
testy.set_obj(test_y_a_{index})
testy.no_mock()

dataset.add_asset(trainx)
dataset.add_asset(trainy)
dataset.add_asset(testx)
dataset.add_asset(testy)

alice_do_{index}.upload_dataset(dataset)  # Aggiunta dell'indice alla chiamata del metodo upload_dataset
"""

# Creazione della directory per gli script
if not os.path.exists("scripts"):
    os.makedirs("scripts")

# Generazione degli script
for i in range(num_scripts):
    script_filename = f"scripts/script_{i+1}.py"
    with open(script_filename, "w") as script_file:
        script_content_formatted = script_content.format(index=i+1)
        script_file.write(script_content_formatted)

print(f"{num_scripts} script creati con successo.")

