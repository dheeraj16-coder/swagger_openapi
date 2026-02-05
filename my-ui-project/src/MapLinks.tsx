import type { MapsDetails } from './api-client';

function MapLinks({ maps }: { maps: MapsDetails }) {
  return (
    <>
      <a href={maps.googleMaps} target="_blank" rel="noopener noreferrer">
        Google Maps
      </a>
      {' | '}
      <a href={maps.openStreetMaps} target="_blank" rel="noopener noreferrer">
        OpenStreetMap
      </a>
    </>
  );
}

export default MapLinks;