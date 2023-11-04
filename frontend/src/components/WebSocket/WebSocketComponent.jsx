import { useEffect } from 'react';

import constants from '../../assets/constants';
import { useParkingLotManager } from '../../hooks/useParkingLotManager';

function WebSocketComponent() {
  const { updateParkingLotState } = useParkingLotManager();

  const ws = new WebSocket(constants.SOCKET_URL);

  useEffect(() => {
    ws.onmessage = (event) => {
      updateParkingLotState(event.data);
    };

    ws.onclose = () => {
    };

    return () => {
      ws.close();
    };
  });

  return null;
}

export default WebSocketComponent;
