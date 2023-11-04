import { useRecoilState } from 'recoil';

export function useParkingLotManager() {
  const [occupiedParkingLots, setOccupiedParkingLots] = useRecoilState([]);

  const updateParkingLotState = (message) => {
    if (message.isOccupied) {
      setOccupiedParkingLots((prevOccupiedLots) => [...prevOccupiedLots, message.id]);
    } else if (!message.isOccupied) {
      setOccupiedParkingLots((prevOccupiedLots) =>
        prevOccupiedLots.filter((lotId) => lotId !== message.id)
      );
    }
  };

  return {
    occupiedParkingLots,
    updateParkingLotState,
  };
}
