import warnings
warnings.filterwarnings("ignore")

import numpy as np

def federated_averaging(global_model, participant_models):
    """
    Aggiorna il modello globale eseguendo la media federata dei pesi dei modelli partecipanti.

    :param global_model: Il modello globale da aggiornare.
    :param participant_models: Una lista di modelli partecipanti.
    :return: Il modello globale aggiornato.
    """
    # Numero di partecipanti
    num_participants = len(participant_models)
    if num_participants == 0:
        raise ValueError("Non ci sono modelli partecipanti per l'averaging.")

    # Estrai i pesi dal primo modello per inizializzare la struttura della somma
    sum_weights = [np.zeros_like(weights) for weights in global_model.get_weights()]

    # Somma tutti i pesi dei modelli partecipanti
    for participant_model in participant_models:
        participant_weights = participant_model.get_weights()
        for i, layer_weights in enumerate(participant_weights):
            sum_weights[i] += layer_weights

    # Calcola la media dei pesi e aggiorna il modello globale
    averaged_weights = [sum_weight / num_participants for sum_weight in sum_weights]
    global_model.set_weights(averaged_weights)
    print(averaged_weights)

    return global_model