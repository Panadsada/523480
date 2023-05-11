import { Alert, Box, Divider, FormControl, Grid, Select, Snackbar, TextField, Typography } from "@mui/material";
import React, { useEffect, useState } from "react";
import Button from "@mui/material/Button";
// import {Link as RouterLink} from "react-router-dom"
// import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';



  function DrugCreate(){
    const [drugCode, setDrugCode] = useState<string>();
    const [drugName, setDrugName] = useState<string>();
    const [amount, setAmount] = useState<any>();
    const [price, setPrice] = useState<any>();
    const [drugUnit, setDrugUnit] = useState<any>();

    const submit = () => {
      console.log("SUBMIT")
      let data = {
        durgCode: drugCode,
        drugName: drugName,
        amount: amount,
        price: price,
        unit: drugUnit
      }
      console.log("DATA: ", data)
    }

    const unit =[
      {id:1, label: "กรัม"},
      {id:2, label: "กิโลกรัม"}
    ]
    return(
        <Box sx={{
            backgroundImage:
              "url(https://i.pinimg.com/564x/c0/4a/5f/c04a5f27b605cc21aa55420a7f83318d.jpg)",
            backgroundRepeat: "no-repeat",
            backgroundSize: "cover",
            width: "100%",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}>
      
            <Box
              sx={{
                fontFamily: "PK Krung Thep Medium",
                fontSize: "20px",
                width: "90%",
                mt: "70px",
                bgcolor: 'rgba(255, 255, 255, 0.5)',
                borderRadius: "30px",
                boxShadow: 20,
      
              }}
            >
      
              <Snackbar open={false} autoHideDuration={6000} onClose={()=>{}}>
                <Alert onClose={() => {}} severity="success">
                  บันทึกข้อมูลสำเร็จ
                </Alert>
              </Snackbar>
              <Snackbar open={false} autoHideDuration={6000} onClose={() => {}}>
                <Alert onClose={() => {}} severity="error">
                  บันทึกข้อมูลไม่สำเร็จ: {}
                </Alert>
              </Snackbar>
      
              <Box sx={{ padding: 2, color: "text.secondary" }}>
                <Box display="flex">
                  <Box flexGrow={1}>
                    <Typography
                      component="h6"
                      variant="h5"
                      color="primary"
                      gutterBottom
                      sx={{
                        display: 'flex',
                        justifyContent: 'center',
                        
                      }}
                    >
                      <h1>เพิ่มข้อมูลยา</h1>
      
                    </Typography>
                  </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ flexGrow: 1 }}>
      
                  <Grid item xs={4}>
                    <FormControl fullWidth variant="outlined">
                      <p>รหัสยา</p>
                      <TextField
                        // variant="outlined"
                        id="drugCode"
                        name="drugCode"
                        placeholder="Drug code"
                        value={drugCode}
                        onChange={(e) => {setDrugCode(e.target.value)}}
                      />
                    </FormControl>
                  </Grid>

                  <Grid item xs={4}>
                    <FormControl fullWidth variant="outlined">
                      <p>ชื่อยา</p>
                      <TextField
                        // variant="outlined"
                        id="drugName"
                        name="drugName"
                        placeholder="Drug name"
                        value={drugName}
                        onChange={(e) => {setDrugName(e.target.value)}}
                      />
                    </FormControl>
                  </Grid>

                  <Grid item xs={4}>
                    <FormControl fullWidth variant="outlined">
                      <p>หน่วยสินค้า</p>
                      <Select
                        native
                        value={drugUnit}
                        onChange={(e) => {setDrugUnit(e.target.value)}}
                        name="unit"
                        placeholder="--กรุณาเลือก--"
                      >
                        <option aria-label="None" value="">
                          --กรุณาเลือก--
                        </option>
                        {unit.map((item, index) => (
                          <option value={item.id} key={index}>
                            {item.label}
                          </option>
                        ))}
                      </Select>
                    </FormControl>
                  </Grid>

                  <Grid item xs={4}>
                    <FormControl fullWidth variant="outlined">
                      <p>จำนวนสินค้า</p>
                      <TextField
                        // variant="outlined"
                        id="amount"
                        name="amount"
                        placeholder="Amount"
                        inputMode="numeric"
                        value={amount}
                        onChange={(e) => {setAmount(e.target.value)}}
                      />
                    </FormControl>
                  </Grid>
      
                  <Grid item xs={4}>
                    <FormControl fullWidth variant="outlined">
                      <p>ราคา</p>
                      <TextField
                        // variant="outlined"
                        id="price"
                        name="price"
                        placeholder="Price"
                        value={price}
                        onChange={(e) => {setPrice(e.target.value)}}
                      />
                    </FormControl>
                  </Grid>
                  {/* <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <p>ธนาคาร</p>
                      <Select
                        native
                        value={[]}
                        onChange={() => {}}
                        inputProps={{
                          name: "BankingID",
                        }}
                      >
                        <option aria-label="None" value="">
                          เลือก...
                        </option>
                        {bankings.map((item: BankingInterface) => (
                          <option value={item.ID} key={item.ID}>
                            {item.Name}
                          </option>
                        ))}
                      </Select>
                    </FormControl>
                  </Grid> */}
      
      
                  {/* <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <LocalizationProvider dateAdapter={AdapterDateFns}>
                        <p>ช่องทางการชำระ</p>
                        <DateTimePicker
                          value={selectedDate}
                          onChange={(newValue) => setSelectedDate(newValue)}
                          minDate={(new Date)}   //บันทึกค่าปัจจุบัน
                          renderInput={(params) =>
                            <TextField {...params} />}
                        />
                      </LocalizationProvider>
                    </FormControl>
                  </Grid> */}
      
                  <Grid item xs={12} style={{"textAlign": "right"}}>
                    <Button
                    //   component={RouterLink}
                    //   to="/payments"
                      variant="outlined"
                    >
                      กลับ
                    </Button>
                    <Button
                      onClick={submit}
                      variant="outlined"
                    >
                      บันทึกข้อมูล
                    </Button>
                  </Grid>
                </Grid>
              </Box>
            </Box>
          </Box>
    )
  }
  export default DrugCreate