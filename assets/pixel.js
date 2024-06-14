var slider;
var output;
window.onload=function(){
    slider = document.getElementById("BlockinessSlider");
    output = document.getElementById("Blockinessno");
    output.innerHTML = slider.value; 
    slider.oninput = function() {
        output.innerHTML = this.value;
      };
      
}

var cuspalarr=[];
var allpals = new Map();
var nocuspal=1;
var famouspal=["Nintendo","Gameboy","NES"];
var clickedcolor="initial";
var imgtag = document.createElement("img");
var origimg;
var imageuploaded=false;

function addPalette(){

    if(cuspalarr.length>0){
        var palname= document.getElementById("palettename").value.trim();
        if (palname===""){
            palname="Custom "+nocuspal;
        }
    
    if (!(allpals.has(palname))){
        
        selectpal=document.getElementById("palette")
        var opt = document.createElement('option');
        opt.value = nocuspal;
        opt.id = palname;
        opt.innerHTML = palname;
        selectpal.appendChild(opt); 
        nocuspal++;  
        if (document.getElementById("palettename").value.trim()===""){
            document.getElementById("palettename").value=palname;
        }     
    }
    
    allpals.set(palname,cuspalarr.slice());

    }
    
    
}

function addColor(){
    
    var theColor = document.getElementById("cuscolor").value;
    if (!(cuspalarr.includes(theColor))){
        document.getElementById('addedPalette').innerHTML += '<svg width="30" height="30"><rect width="30" height="30" fill="'+theColor+'" /></svg>';
        cuspalarr.push(theColor);
    }
    
}

function paletteDisplay(){
    var selectid = document.getElementById("palette");
    var options = selectid.options;
    var addpal= document.getElementById('addedPalette');
    addpal.innerHTML="";
    var id = options[options.selectedIndex].id;
    
    if (allpals.has(id)){
        cuspalarr=allpals.get(id);        
        document.getElementById("palettename").value=id;
        for (let i = 0; i < cuspalarr.length; i++) {
            addpal.innerHTML+='<svg width="30" height="30"><rect width="30" height="30" fill="'+cuspalarr[i]+'" /></svg>';
        }
    }
    else{
        document.getElementById("palettename").value="";
        cuspalarr=[];
    }

}

function SVGClick(ele){
   
    if (!(clickedcolor==="initial")){
       var svgid = document.getElementById(clickedcolor);
       svgid.childNodes[0].setAttribute("stroke","");
       svgid.childNodes[0].setAttribute("stroke-width","0");
    }
    ele.setAttribute("stroke","white");
    ele.setAttribute("stroke-width","5");
    clickedcolor=ele.parentNode.id;
    
}

function replacePalette(){

    if(!(clickedcolor==="initial")){
        var color=document.getElementById("replacecol").value;
        var svgid= document.getElementById(clickedcolor);
        const oldcolor=svgid.childNodes[0].getAttribute("fill");
        svgid.childNodes[0].setAttribute("stroke","");
        svgid.childNodes[0].setAttribute("stroke-width","0");
        svgid.childNodes[0].setAttribute("fill",color);
        clickedcolor="initial";
        var canvas = document.createElement('canvas');
        canvas.width = imgtag.naturalWidth;
        canvas.height = imgtag.naturalHeight;
        var context = canvas.getContext('2d');
        context.drawImage(imgtag, 0, 0);
        var imgd=context.getImageData(0, 0, canvas.width, canvas.height);        
        var pixdata= imgd.data;
        const r=parseInt(color.slice(1, 3), 16);
        const g=parseInt(color.slice(3, 5), 16);
        const b=parseInt(color.slice(5, 7), 16);
        const oldr=parseInt(oldcolor.slice(1, 3), 16);
        const oldg=parseInt(oldcolor.slice(3, 5), 16);
        const oldb=parseInt(oldcolor.slice(5, 7), 16);

        for(var i=0;i<pixdata.length;i+=4){
            if(pixdata[i]==oldr && pixdata[i+1]==oldg && pixdata[i+2]==oldb)
                {
                    pixdata[i]=r;
                    pixdata[i+1]=g;
                    pixdata[i+2]=b;
                }
        }
        context.putImageData(imgd,0,0);
        imgtag.src=canvas.toDataURL();
    }
    
}


function displayUploaded(event) {
    
    var selectedFile = event.target.files[0];
    var reader = new FileReader();
  
    var divimg= document.getElementById("centered");
    
  
    reader.onload = function(event) {
        imgtag.title = selectedFile.name;
        imgtag.src = event.target.result;
        origimg=event.target.result;
        imgtag.id="displayimage";        
        imgtag.style.maxWidth="100%";
        imgtag.style.maxHeight="91vh";
        imgtag.style.height="auto";       
        divimg.appendChild(imgtag);  
        imageuploaded=true;
        
    };
  
    reader.readAsDataURL(selectedFile);
   
  }

function hideDivsCustom(){
    var cusdiv = document.getElementById("cuspalcheck");
    var reddiv = document.getElementById("redcolcheck");
    
    if (cusdiv.checked){
        document.getElementById("hidedivcus").style.display="block";
        if (reddiv.checked){
            reddiv.checked=false;
            document.getElementById("hidedivred").style.display="none";
        }

    }
    
    if (cusdiv.checked===false){
        document.getElementById("hidedivcus").style.display="none";
    }
    
}

function hideDivsReduce(){
    var cusdiv = document.getElementById("cuspalcheck");
    var reddiv = document.getElementById("redcolcheck");
    if(reddiv.checked){
        document.getElementById("hidedivred").style.display="block";
        if (cusdiv.checked){
            cusdiv.checked=false;
            document.getElementById("hidedivcus").style.display="none";
        }
    }
    if (reddiv.checked===false){
        document.getElementById("hidedivred").style.display="none";
    }
}

async function sendData(){
    
    if(imageuploaded===false){
        return;
    }
   
    var img = "";
    var blockiness = "";
    var grayscale = "" ;
    var dither = "" ;
    var custompalettecheck = "";
    var options = "";    
    var palettename ="" ;
    var reducecolorcheck = "" ;
    var noofColors ="" ;
    var palettecolors = [];
    
    img = origimg;
    blockiness = document.getElementById("BlockinessSlider").value;
    grayscale = document.getElementById("grayscale").checked;
    dither = document.getElementById("dither").value;
    custompalettecheck = document.getElementById("cuspalcheck").checked;
    options = document.getElementById("palette").options;    
    palettename = options[options.selectedIndex].id;
    reducecolorcheck = document.getElementById("redcolcheck").checked;
    noofColors = document.getElementById("redcol").value;


    if (allpals.has(palettename))
    {
        palettecolors = allpals.get(palettename);
    }
     
    const obj ={image:img,block:blockiness,gray:grayscale,dithering:dither,
        customcheck:custompalettecheck,reducecheck:reducecolorcheck,
        palname:palettename,noofcolors:noofColors,
        colorpallete:palettecolors
        };
    const myJSON = JSON.stringify(obj);
    // console.log(myJSON);
    
       var  recvJSON=await fetch("/", {
            
            "method": "POST",
            
            "body": myJSON,
            
            "headers": {
                "Content-type": "application/json"
            }
        });
    const data = await recvJSON.json();    
    document.getElementById("displayimage").src="data:image/png;base64, "+data["Image"];
    
    if(data.hasOwnProperty("Colorpallete")){
        document.getElementById("replaceSVGdiv").innerHTML="";
        
        for (let i=0;i<data["Colorpallete"].length;i++){
            
            var svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");
            svg.setAttribute('width', '30');
            svg.setAttribute('height', '30');
            svg.setAttribute('id', "svgr"+(i+1));
            var rect = document.createElementNS("http://www.w3.org/2000/svg", 'rect');
            rect.setAttribute('height', '30');
            rect.setAttribute('width', '30');
            rect.setAttribute('fill',data["Colorpallete"][i]);
            rect.setAttribute('onclick',"SVGClick(this)");
            svg.appendChild(rect);
            document.getElementById("replaceSVGdiv").appendChild(svg);
        }
    }

}

function downloadData(){
    if(imageuploaded){
        var a =document.getElementById("downloadaimage");
    a.setAttribute("download",imgtag.title.substring(0,imgtag.title.lastIndexOf("."))+"Retro.png");
    a.setAttribute("href",imgtag.src);
    }
    
}