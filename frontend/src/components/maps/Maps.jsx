import React, { useState } from "react";
import { GoogleMap, OverlayView, useJsApiLoader } from "@react-google-maps/api";
import { useRecoilValue } from "recoil";
import { FaTimes } from "react-icons/fa";

import "./Maps.scss";
import { parkingLotsState } from "../../store/state";
import axios from "axios";
import constants from "../../assets/constants";

const containerStyle = {
  width: "100%",
  height: "90%",
};

const center = {
  lat: 45.7972,
  lng: 15.97176,
};

function Map() {
  const { isLoaded } = useJsApiLoader({
    id: "google-map-script",
    googleMapsApiKey: import.meta.env.VITE_GOOGLE_MAPS_API_KEY,
  });

  const parkingLots = useRecoilValue(parkingLotsState);

  const [map, setMap] = useState(null);
  const [selected, setSelected] = useState(null);
  const [timeValue, setTimeValue] = useState("");

  const onNodeClick = (parkingLot) => {
    setSelected(parkingLot);
  };

  const onReserve = (e) => {
    e.preventDefault();

    axios
      .post(
        `${constants.API_URL}/api/reservation/add`,
        JSON.stringify({
          parkingspotid: selected.id,
          endh: parseInt(timeValue.split(" ")[0]),
          endm: parseInt(timeValue.split(" ")[1]),
        })
      )
      .then((res) => {
        console.log(res);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  const onLoad = React.useCallback(function callback(map) {
    map.setZoom(16);

    setMap(map);
  }, []);

  const onUnmount = React.useCallback(function callback(map) {
    setMap(null);
  }, []);

  const handleTimeChange = (event) => {
    setTimeValue(event.target.value);
  };

  return (
    <>
      {isLoaded ? (
        <GoogleMap
          mapContainerStyle={containerStyle}
          center={center}
          onLoad={onLoad}
          onUnmount={onUnmount}
        >
          {parkingLots.data.map((parkingLot) => (
            <OverlayView
              key={parkingLot.id}
              position={{ lat: parkingLot.latitude, lng: parkingLot.longitude }}
              mapPaneName={OverlayView.OVERLAY_MOUSE_TARGET}
            >
              <div
                className={`map-marker ${
                  parkingLot.occupied ? "occupied" : "free"
                } ${
                  selected && parkingLot.id === selected.id ? "selected" : ""
                }`}
                onClick={() => onNodeClick(parkingLot)}
              />
            </OverlayView>
          ))}
        </GoogleMap>
      ) : null}
      {selected && (
        <div className="modal">
          <FaTimes className="modal-icon" onClick={() => setSelected(null)} />
          <div className="modal-title">
            Parking spot in {selected.parkingSpotZone}
          </div>
          <div className="modal-data">
            <div className="modal-data-entry">
              Currently{" "}
              <span
                className={selected.occupied ? "occupied-text" : "free-text"}
              >
                {selected.occupied ? "occupied" : "free"}
              </span>
            </div>
            {selected.occupied && (
              <div className="modal-data-entry">
                Occupied until{" "}
                {new Date(selected.occupiedTimestamp).toLocaleTimeString()}
              </div>
            )}
            <div className="modal-data-entry modal-data-weak">
              Position: lat - {selected.latitude} : lng - {selected.longitude}
            </div>
            {!selected.occupied && (
              <form>
                <input
                  type="time"
                  value={timeValue}
                  onChange={handleTimeChange}
                ></input>
                <input
                  type="submit"
                  value="Reserve"
                  onClick={onReserve}
                ></input>
              </form>
            )}
          </div>
        </div>
      )}
    </>
  );
}

export default Map;
