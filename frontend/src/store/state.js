import axios from 'axios';
import { atom } from 'recoil';

import constants from '../assets/constants';

export const parkingLotsState = atom({
  key: 'parkingLots',
  default: axios.get(`${constants.API_URL}/api/ParkingSpot/getAll`),
});

export const occupiedParkingLotsState = atom({
  key: 'occupiedParkingLots',
  default: [],
});
