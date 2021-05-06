import React from 'react';
import {Inject, ScheduleComponent, Week, Month} from "@syncfusion/ej2-react-schedule";
import {ResourceDirective} from "@syncfusion/ej2-react-schedule/src/schedule/resources-directive";

const Calendar = () => {

    let localData = {
        dataSource: [{
            Id: 1,
            EndTime: new Date(2021,4, 6,15),
            StartTime: new Date(2021,4,4, 7,15),
            Subject: 'TITULO',

            ResourceId: 1,
        }],
        /*fields: {
            endTime: {name: 'End'},
            startTime: {name: 'Start'},
            subject: {name: 'Summary'},
        }*/
    }

    let resourceDataSource= [
        {Name: 'cris', Id:1, Color: '#357CD2'},
        {Name: 'alex', Id:2, Color: '#FFF'},
    ]

    return (
        <ScheduleComponent
            height = '650px'
            width = '1400px'
            currentView='Month'
            eventSettings={localData}
        >

            <Inject services={[Week,Month]}/>
        </ScheduleComponent>
    );
};

export default Calendar;