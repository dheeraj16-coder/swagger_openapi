import { CarDetails } from './api-client';
import React from 'react';

// This is a TypeScript interface that describes the "shape" of the car object.
interface CarDetailsProps {
 car: CarDetails | undefined;
}

// The component receives one prop: the car object.
function CarDetailsDisplay({ car }: CarDetailsProps) {
  if (!car) {
    return <span>No car details available</span>;
  }

  return (
    <span>
      car signs = {car.signs ? car.signs.join(', ') : 'No signs available'}
      

      {' and '}
      
      driver side = {car.side}
    </span>
  );
}

export default CarDetailsDisplay;