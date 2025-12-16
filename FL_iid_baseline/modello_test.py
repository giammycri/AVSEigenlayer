import warnings
warnings.filterwarnings("ignore")

from tensorflow.keras import layers, models

    # Definisci il modello CNN
modello_test = models.Sequential([
    layers.Conv2D(16, (1, 1), activation='relu', input_shape=(28, 28, 1)),
    #layers.MaxPooling2D((2, 2)),
    #layers.Conv2D(16, (1, 1), activation='relu'),
    # layers.MaxPooling2D((2, 2)),
    # layers.Conv2D(64, (3, 3), activation='relu'),
    layers.Flatten(),
    #layers.Dense(64, activation='relu'),
    layers.Dense(10, activation='softmax')
])

    # Compila il modello
modello_test.compile(optimizer='adam',
            loss='sparse_categorical_crossentropy',
            metrics=['accuracy'])

#modello_test.summary();