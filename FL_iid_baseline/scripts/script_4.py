
import syft as sy
from keras.datasets import mnist
from syft.service.user.user import ServiceRole
from syft.service.user.user import UserCreate

domain_client = sy.login(port="3331", email="info@openmined.org", password="changethis")
domain_client.users.create(UserCreate(email="alice4@openmined.org", name="Alice4", password="pw", role=ServiceRole.DATA_OWNER))
alice_do_4 = domain_client.login(email="alice4@openmined.org", password="pw")

(train_X_a_4, train_y_a_4), (test_X_a_4, test_y_a_4) = mnist.load_data()

dataset = sy.Dataset(name="MNIST_4")

trainx = sy.Asset(name="Asset di Train")
trainx.set_obj(train_X_a_4)
trainx.no_mock()

trainy = sy.Asset(name="Asset di etichette di Train")
trainy.set_obj(train_y_a_4)
trainy.no_mock()

testx = sy.Asset(name="Asset di Test")
testx.set_obj(test_X_a_4)
testx.no_mock()

testy = sy.Asset(name="Asset di eitchette di Test")
testy.set_obj(test_y_a_4)
testy.no_mock()

dataset.add_asset(trainx)
dataset.add_asset(trainy)
dataset.add_asset(testx)
dataset.add_asset(testy)

alice_do_4.upload_dataset(dataset)  # Aggiunta dell'indice alla chiamata del metodo upload_dataset
